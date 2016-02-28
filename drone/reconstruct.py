#! /usr/local/bin/python

"""
Generic References:
http://www.codinglabs.net/article_world_view_projection_matrix.aspx
"""

import argparse
import math
import os

import numpy as np
from PIL import Image


def parse_metadata(metadata):
    """Parse metadata from a file-like object"""
    data = {}
    for line in metadata:
        if line.startswith('#'):
            continue
        filename, x, y, z, yaw, pitch, roll = line.strip().split(',')
        data[filename] = {'lat': float(x),
                          'lon': float(y),
                          'alt': float(z),
                          'yaw': float(yaw),
                          'pitch': float(pitch),
                          'roll': float(roll)}
    return data


def get_center(img_metadata):
    """Return the center of the images"""
    lat_lon_alt = []
    for img, metadata in img_metadata.iteritems():
        lat_lon_alt.append([metadata['lat'],
                            metadata['lon'],
                            metadata['alt']])
    lat_lon_alt = np.array(lat_lon_alt)
    return np.mean(lat_lon_alt, axis=0)


def convert_to_ecef(lat, lon, alt):
    """
    Return lat/lon/alt as ECEF coordinates

    References:
    https://en.wikipedia.org/wiki/ECEF
    https://microem.ru/files/2012/08/GPS.G1-X-00006.pdf
    """
    # Define WGS84 Constants
    f = 1 / 298.257223563
    a = 6378137  # Radius of earth in meters
    b = a * (1 - f)
    e = math.sqrt((a**2 - b**2) / a**2)

    # Convert lat/lon from degrees to radians
    lat_rad = lat * math.pi / 180
    lon_rad = lon * math.pi / 180

    # Pre-calculate cos and sin values
    cos_lat = math.cos(lat_rad)
    sin_lat = math.sin(lat_rad)
    cos_lon = math.cos(lon_rad)
    sin_lon = math.sin(lon_rad)

    # Calculate Radius of Curvature
    N = a / math.sqrt(1 - e**2 * sin_lat**2)
    x = (N + alt) * cos_lat * cos_lon
    y = (N + alt) * cos_lat * sin_lon
    z = ((b**2/a**2) * N + alt) * sin_lat
    return x, y, z


def convert_to_ned(x_y_z, center, center_lat_lon_alt):
    """
    Convert ECEF coordinates to a North-East-Down system centered on the
    average center of the different pictures.  Both sets have to be in
    ECEF.

    Reference:
    https://en.wikipedia.org/wiki/North_east_down
    https://microem.ru/files/2012/08/GPS.G1-X-00006.pdf
    """
    delta = np.array(x_y_z) - np.array(center)

    # Pre-calculate cos and sin values for center
    center_lat_rad = center_lat_lon_alt[0] * math.pi / 180
    center_lon_rad = center_lat_lon_alt[2] * math.pi / 180
    cos_lat = math.cos(center_lat_rad)
    sin_lat = math.sin(center_lat_rad)
    cos_lon = math.cos(center_lon_rad)
    sin_lon = math.sin(center_lon_rad)

    ned = np.asmatrix([[-sin_lat * cos_lon, -sin_lat * sin_lon, cos_lat],
                       [-sin_lon, cos_lon, 0],
                       [cos_lat * cos_lon, cos_lat * sin_lon, sin_lat]])
    return ned * delta


def scale_matrix(focal_length):
    """
    Return a scale matrix given the focal length of the camera.

    Reference:
    https://en.wikipedia.org/wiki/Camera_matrix
    """
    # Convert focal length to meters
    fl = focal_length / 1000.0
    scale = np.asmatrix([[fl, 0, 0, 0],
                         [0, fl, 0, 0],
                         [0, 0, 1, 0],
                         [0, 0, 0, 1]])
    return scale


def rotation_matrix(roll, pitch, yaw):
    """
    Return a rotation matrix given the roll, pitch, and yaw in degrees.

    Reference:
    https://en.wikipedia.org/wiki/Rotation_matrix#In_three_dimensions
    """

    # Convert to radians
    roll = roll * math.pi / 180
    pitch = pitch * math.pi / 180
    yaw = yaw * math.pi / 180

    # roll rotation
    rx = np.asmatrix([[1, 0, 0, 0],
                      [0, math.cos(roll), -math.sin(roll), 0],
                      [0, math.sin(roll), math.cos(roll), 0],
                      [0, 0, 0, 1]])

    # pitch rotation
    ry = np.asmatrix([[math.cos(pitch), 0, math.sin(pitch), 0],
                      [0, 1, 0, 0],
                      [-math.sin(pitch), 0, math.cos(pitch), 0],
                      [0, 0, 0, 1]])

    # yaw rotation
    rz = np.asmatrix([[math.cos(yaw), -math.sin(yaw), 0, 0],
                      [math.sin(yaw), math.cos(yaw), 0, 0],
                      [0, 0, 1, 0],
                      [0, 0, 0, 1]])

    # Construct full rotation matrix
    rotation = rx * ry * rz
    return rotation


def translation_matrix(lat, lon, alt, center, center_lat_lon_alt):
    """
    Return a translation matrix given the x, y, and z in meters and the
    local center coordinates.
    """
    x_y_z = convert_to_ecef(lat, lon, alt)
    x, y, z = convert_to_ned(x_y_z, center, center_lat_lon_alt)
    return np.asmatrix([[1, 0, 0, -x],
                        [0, 1, 0, -y],
                        [0, 0, 1, -z],
                        [0, 0, 0, 1]])


def mosaic(img_dir, img_metadata, focal_length):

    # Get the height and width of images
    img_name = sorted(img_metadata.keys())[0]
    img_path = os.path.join(img_dir, img_name)
    img_height, img_width = Image.open(img_path).size

    # Get the center of the images
    center_lat_lon_alt = get_center(img_metadata)
    center = convert_to_ecef(*center_lat_lon_alt)

    scale = scale_matrix(focal_length)
    for img, metadata in img_metadata.iteritems():
        rotation = rotation_matrix(metadata['roll'],
                                   metadata['pitch'],
                                   metadata['yaw'])

        translation = translation_matrix(metadata['lat'],
                                         metadata['lon'],
                                         metadata['alt'],
                                         center,
                                         center_lat_lon_alt)
        projection = scale * rotation * translation
        print projection


def main(args):
    # Get the image directory
    img_dir = os.path.abspath(os.path.dirname(args.metadata.name))

    # Parse out the metadata file for all images
    img_metadata = parse_metadata(args.metadata)

    # Mosaic the images
    mosaic(img_dir, img_metadata, args.focal_length)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Reconstruct Images.')
    parser.add_argument('metadata', type=argparse.FileType('r'),
                        help='The file containing image metadata')
    parser.add_argument('-c', '--format', default=35, type=int,
                        help='The format of the camera in mm')
    parser.add_argument('-l', '--focal_length', default=20, type=int,
                        help='The focal length of the camera in mm')
    args = parser.parse_args()
    main(args)

#! /usr/bin/env python

import math


def length(f1, a, Y, d):
    # http://hyperphysics.phy-astr.gsu.edu/hbase/Music/barres.html
    # f1 primary frequency (Hz)
    # a thickness (m)
    # L length (m)
    # Y young's modules (GPa)
    # d density (kg/m3)

    # Convert GPa to Pa
    Y *= math.pow(10, 9)
    L = math.sqrt(0.162 * (a / f1) * math.sqrt(Y / d))
    return L * 100  # m -> cm


if __name__ == "__main__":
    # https://www.wood-database.com
    woods = {
        "sweet cherry": {
            "Y": 10.55,  # Young's (Elastic) Modulus (GPa)
            "d": 600,  # Density (kg/m3)
        },
        "african padauk": {
            "Y": 11.72,
            "d": 745,
        },
        "cocobolo": {
            "Y": 18.70,
            "d": 1095,
        },
        "canarywood": {
            "Y": 14.93,
            "d": 830,
        },
        "pau rosa": {
            "Y": 17.10,
            "d": 1030,
        },
    }

    chords = {
        # I-IV-V-vi
        "C": [261.626, 329.628, 381.995],  # I
        "F": [440.000, 523.251, 659.255],  # VI
        "G": [391.995, 493.883, 587.330],  # V
        "Am": [349.288, 440.000, 523.251],  # vi
    }
    # Thickness (m)
    inch_to_meter = 0.0254
    # a = 1 / 2
    # a = 5 / 8
    a = 3 / 4

    for wood, vals in woods.items():
        Y = vals["Y"]
        d = vals["d"]
        print(f"{wood} @ {a} in - Y: {Y:.2f} GPa, d: {d} kg/m3\n")
        for chord, freq_list in chords.items():
            print(f"{chord} chord")
            for i, freq in enumerate(freq_list):
                L = length(freq, a * inch_to_meter, Y, d)
                print(f"\t{2*i+1}: {L:.2f} cm or {L*0.393701:.2f} in")
        print()

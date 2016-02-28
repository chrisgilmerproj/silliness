Reconstruction Challenge
===

Instructions
---

This zip file (https://s3.amazonaws.com/drone.deploy.map.engine/example.zip) contains 24 images of a
massive rock near the beautiful Goat Rock Beach in Sonoma, CA. These images can be used to
create a 3D reconstruction of the rock which looks like this:
https://sketchfab.com/models/58a312fb79b94867812ecd1f42cd053f
We would like you to code the first stage of this reconstruction which involves placing the images
correctly to create a mosaic. Use the information below about the locations of the cameras to draw
all these images in the correct location on the plane. For example:

TBD INSERT IMAGES

The camera used to take these images had a 35mm focal length of 20mm.
The follow gives the locations and pose of each image in the zip file.

```csv
# Filename,X,Y,Z,Yaw,Pitch,Roll
dji_0644.jpg,­123.114661,38.426805,90.689292,9.367337,1.260910,0.385252
dji_0645.jpg,­123.114650,38.426831,90.825989,85.055542,­0.336052,1.667057
dji_0646.jpg,­123.114429,38.426830,91.088004,88.858391,­0.070967,1.876991
dji_0647.jpg,­123.114125,38.426831,91.091265,88.269956,0.671020,1.849037
dji_0648.jpg,­123.114104,38.426832,90.747063,184.433167,­1.492852,1.134858
dji_0649.jpg,­123.114136,38.426609,91.304548,190.422786,­0.656365,1.312138
dji_0650.jpg,­123.114203,38.426195,91.007241,190.053859,0.363708,1.444969
dji_0651.jpg,­123.114271,38.425813,91.538639,190.037347,1.106723,1.521566
dji_0652.jpg,­123.114284,38.425752,90.900331,190.344637,1.424554,1.632872
dji_0653.jpg,­123.114268,38.425751,90.622088,89.052669,1.243665,­1.090830
dji_0654.jpg,­123.113839,38.425752,91.235595,88.392906,1.794960,­0.221090
dji_0655.jpg,­123.113745,38.425749,90.437221,87.186642,1.947206,0.394757
dji_0656.jpg,­123.113734,38.425779,90.163445,6.838638,0.624994,­0.674300
dji_0657.jpg,­123.113662,38.426160,91.160272,6.815734,0.945930,0.550999
dji_0658.jpg,­123.113591,38.426581,91.454023,8.740611,1.059218,1.088282
dji_0659.jpg,­123.113556,38.426807,91.221973,9.253228,1.353285,1.449262
dji_0660.jpg,­123.113544,38.426829,90.324952,146.612422,­1.948292,0.194904
dji_0661.jpg,­123.113439,38.426665,90.864808,155.415639,­0.917097,1.375369
dji_0662.jpg,­123.113183,38.426287,91.956351,155.074334,0.208305,2.160615
dji_0663.jpg,­123.113116,38.426189,90.561950,153.763228,0.793427,2.490934
dji_0664.jpg,­123.113115,38.426165,90.604094,187.491139,­0.312975,2.836182
dji_0665.jpg,­123.113176,38.425826,91.781148,188.845376,0.574889,3.010090
dji_0666.jpg,­123.113185,38.425756,91.069673,189.163989,0.764728,2.785707
dji_0667.jpg,­123.113198,38.425754,90.750004,301.431548,­2.034127,0.511803
```

Code should be written in Python and you are free to use any libraries you choose. The final code should be
created in a GitHub repository with a README.md explaining how to run it locally. Have fun!
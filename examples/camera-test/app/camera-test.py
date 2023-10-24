# SPDX-FileCopyrightText: Copyright 2023 LG Electronics Inc.
# SPDX-License-Identifier: Apache-2.0

# -*- coding: utf-8 -*-
import cv2
import face_recognition

CAM_ID = 0

cam = cv2.VideoCapture(CAM_ID) # generate camera
if cam.isOpened() == False: # check camera
    print ('Can\'t open the CAM(%d)' % (CAM_ID))
    exit()

# get camera resolution
width = cam.get(cv2.CAP_PROP_FRAME_WIDTH);
height = cam.get(cv2.CAP_PROP_FRAME_HEIGHT);
print ('size = [%f, %f]\n' % (width, height))

# make window and change size
cv2.namedWindow('CAM_Window')
cv2.resizeWindow('CAM_Window', 1280, 720)

while(True):
    # get image frame
    ret, frame = cam.read()

#    small_frame = cv2.resize(frame, (0, 0), fx=0.25, fy=0.25)
#    rgb_small_frame = small_frame[:, :, ::-1]
    face_locations = face_recognition.face_locations(frame)

    if len(face_locations) > 0:
        cv2.rectangle(frame, (face_locations[0][1], face_locations[0][0]), (face_locations[0][3],face_locations[0][2]), (0,0,255), 2)


    # display image on window
    cv2.imshow('CAM_Window', frame)

    # wait key 10ms
    if cv2.waitKey(10) >= 0:
        break;

# release
cam.release()
cv2.destroyWindow('CAM_Window')

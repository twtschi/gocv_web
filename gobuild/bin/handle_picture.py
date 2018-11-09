from __future__ import print_function
from utils.objDet_utils import *
import argparse
import os
import cv2

# Path to frozen detection graph. This is the actual model that is used for the object detection.
PATH_TO_CKPT = 'model/frozen_inference_graph.pb'

if __name__ == '__main__':
    # construct the argument parse and parse the arguments
    ap = argparse.ArgumentParser()
    ap.add_argument("-i", "--input", type=str, default="test.jpg",
                help="Input picture to identify")
    ap.add_argument("-o", "--output", type=str, default="output.jpg",
                help="Output picture")
    ap.add_argument('-w', '--num-workers', dest='num_workers', type=int,
                default=2, help='Number of workers.')
    args = vars(ap.parse_args())

    frame = cv2.imread(args['input'])
    # frame_rgb = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)

    # Load a (frozen) Tensorflow model into memory.
    detection_graph = tf.Graph()
    with detection_graph.as_default():
        od_graph_def = tf.GraphDef()
        with tf.gfile.GFile(PATH_TO_CKPT, 'rb') as fid:
            serialized_graph = fid.read()
            od_graph_def.ParseFromString(serialized_graph)
            tf.import_graph_def(od_graph_def, name='')
        sess = tf.Session(graph=detection_graph)

    result = detect_objects(frame, sess, detection_graph)
    cv2.imwrite(args['output'], result)
    sess.close()
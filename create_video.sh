#!/bin/sh
source ~/.bash_profile
FILENAME=`date --date="1 day ago" +%Y%m%d`
VIDEONAME=$FILENAME.mp4
cd ~/earthin24
./frames
ffmpeg -framerate 10 -pattern_type glob -i '*.png' -s:v 1280x720 -c:v libx264 -profile:v high -crf 20 -pix_fmt yuv420p ~/earthin24/$VIDEONAME
# extract first frame as jpg for instagram
ffmpeg -i $VIDEONAME -vf "select=eq(n\,0)" -q:v 3 -s:v 720x720 $FILENAME.jpg
# Code for the below is not in the repo this script will just create an mp4 and a poster jpg
#./tweet $VIDEONAME
#node instagram.js $FILENAME
rm ~/earthin24/*.png
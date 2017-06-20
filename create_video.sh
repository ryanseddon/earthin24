#!/bin/sh
./video
ffmpeg -framerate 10 -i %d.png -s:v 1280x720 -c:v libx264 -profile:v high -crf 20 -pix_fmt yuv420p `date -v-1d +%Y%m%d`.mp4
rm `find . -name '*.png' ! -name 'error.png'`

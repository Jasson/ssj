
rm -rf build/ && make html
scp -r /Users/jason/work/myblog/build/html/  root@39.106.13.51:/opt/mylog/

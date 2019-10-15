FROM alpine
WORKDIR /
ADD ssj /
ADD conf/app_api.ini /conf/app.ini
#为什么出错
ADD static/myblog/build/html /static/myblog/build/html
ADD static/ static/
ENTRYPOINT ["/ssj"]
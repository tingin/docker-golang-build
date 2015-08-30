# GOLANG Build Test
#
# VERSION       1.0

# use the ubuntu base image provided by dotCloud
FROM golang:1.5

RUN wget -q https://dl-ssl.google.com/linux/linux_signing_key.pub  
RUN apt-key add linux_signing_key.pub
RUN echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list 
RUN apt-get update 
RUN apt-get -y install google-chrome-stable  
RUN apt-get -y install xvfb 
RUN apt-get -y install xfonts-base x11-xkb-utils xfonts-cyrillic xfonts-100dpi xfonts-75dpi 
RUN apt-get install unzip
RUN wget -q http://chromedriver.storage.googleapis.com/2.19/chromedriver_linux64.zip
RUN unzip chromedriver_linux64.zip -d /bin
RUN chmod +x /bin/chromedriver
RUN LINUXVERSION=`cat /etc/issue`
RUN CHROMEVERION=`google-chrome -version`
RUN DRIVERVERISON=`chromedriver -v`
RUN mkdir -p /gopath/src/docker-golang-build
RUN mv ./* /gopath/src/docker-golang-build
RUN rm -rf /var/lib/apt/lists/*

ENV LINUX_VERSION  $LINUXVERSION
ENV CHROME_VERSION $CHROMEVERION
ENV DRIVER_VERSION $DRIVERVERISON

CMD ["/bin/sh"]

EXPOSE 8080

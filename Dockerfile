# GOLANG Build Test
#
# VERSION       1.0

# use the ubuntu base image provided by dotCloud
FROM golang:1.5-onbuild

RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | sudo apt-key add - \
	&& echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list \
	&& apt-get update \
	&& apt-get -y install google-chrome-stable  \
  	&& apt-get -y install xvfb \
 	&& apt-get -y install xfonts-base x11-xkb-utils xfonts-cyrillic xfonts-100dpi xfonts-75dpi \
 	&& wget -q http://chromedriver.storage.googleapis.com/2.19/chromedriver_linux64.zip
	&& unzip chromedriver_linux64.zip -d /bin
	&& chmod +x /bin/chromedriver
	&& linux_version = $(cat /etc/issue)
	&& chrome_verion = $(google-chrome -version)
	&& driver_verison = $(chromedriver -v)
	&& mkdir -p /gopath/src/docker-golang-build
	&& mv ./* /gopath/src/docker-golang-build
	&& rm -rf /var/lib/apt/lists/*

ENV LINUX_VERSION  $linux_version
ENV CHROME_VERSION $chrome_verion
ENV DRIVER_VERSION $driver_verison

EXPOSE 8080

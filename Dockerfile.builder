# Ubuntu 18.04 is our base image for building
FROM ubuntu:18.04

# set up timezone
ENV TZ=GMT
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# TODO: is this needed at all?
## install docker
#RUN apt -y update
#RUN apt -y install curl gnupg lsb-release software-properties-common
#RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
#RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
#RUN apt -y install docker-ci docker-ci-cli containerd.io

# update software
RUN apt -y update
RUN apt -y full-upgrade
RUN apt -y autoremove
RUN apt -y clean

# install development tools
RUN apt -y install git vim htop apg jq direnv build-essential wget awscli

# golang variables
ARG golang_version="1.15.8"
ARG golang_hostarch="linux-amd64"
ARG golang_filename="go${golang_version}.${golang_hostarch}.tar.gz"
ARG golang_url="https://golang.org/dl/${golang_filename}"
ARG golang_sha256="d3379c32a90fdf9382166f8f48034c459a8cc433730bc9476d39d9082c94583b"

# install golang
RUN wget -nv ${golang_url}
RUN echo "${golang_sha256} ${golang_filename}" > "${golang_filename}.sha256"
RUN sha256sum -c ${golang_filename}.sha256
RUN tar -C /usr/local -xzf ${golang_filename}
RUN rm -rf ${golang_filename}*
ENV PATH="${PATH}:/usr/local/go/bin"

# nodejs variables
ARG nodejs_version="12.22.1"
ARG nodejs_hostarch="linux-x64"
ARG nodejs_spec="node-v${nodejs_version}-${nodejs_hostarch}"
ARG nodejs_filename="${nodejs_spec}.tar.gz"
ARG nodejs_url="https://nodejs.org/download/release/v12.22.1/${nodejs_filename}"
ARG nodejs_sha256="d315c5dea4d96658164cdb257bd8dbb5e44bdd2a7c1d747841f06515f23a0042"

# install nodejs
RUN wget -nv ${nodejs_url}
RUN echo "${nodejs_sha256} ${nodejs_filename}" > "${nodejs_filename}.sha256"
RUN sha256sum -c ${nodejs_filename}.sha256
RUN tar -C /usr/local -xzf ${nodejs_filename}
RUN rm -rf ${nodejs_filename}*
ENV PATH="${PATH}:/usr/local/${nodejs_spec}/bin"

# install node packages
RUN npm install -g yarn

# clone core node repository and install dependencies
ARG repository_remote="https://github.com/energicryptocurrency/energi3"
RUN mkdir "/builder"
RUN pushd "/builder"
RUN git clone "${repository_remote}"
RUN pushd "energi3"
RUN npm install
RUN make -f Makefile.release release-tools
RUN popd
RUN popd

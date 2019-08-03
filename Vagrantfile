
Vagrant.configure("2") do |config|
    config.vm.define 'builder' do |node|
        node.vm.provider "virtualbox" do |v|
            v.memory = 4096
        end
        node.vm.box = "bento/ubuntu-18.04"

        node.vm.provision "shell", inline: %Q[
            apt-get update;
            apt-get install -y python3-pip unzip;
            pip3 install -U futoin-cid;
            cid tool install docker;
            sudo adduser vagrant docker;
        ]
    end
end

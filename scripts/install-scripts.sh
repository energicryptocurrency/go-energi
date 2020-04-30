#!/usr/bin/env bash

PWD=$(pwd)
CWD=$(cwd)

WALLET_ADDR=""

echo ""
echo "PWD [${PWD}]"
echo ""

while true
do
	echo ""
	echo "Only run this script once to install."
	echo "Do NOT move the files once installed!"
	echo ""

	read -p "Ok? [Y/y/N/n] " ANSWER

	case $ANSWER in

   		[yY]* ) break;;
	   	[nN]* ) exit;;
   		* )     echo "Only enter [Y/y/N/n] please!.";
 	esac

done

while true
do

        echo ""

        if [[ "$WALLET_ADDR" == "WALLET_ADDRESS_HERE" || -z "${WALLET_ADDR}" ]] ; then

                echo "No wallet address is set! You may edit this script or enter it below."

                read -p "Enter wallet address: " WALLET_ADDR

                if [ -n "$WALLET_ADDR" ]; then

                        break

                fi

        fi

done


export PATH=${PATH}:${PWD}/

echo "" >> ~/.bashrc
echo "" >> ~/.bashrc
echo "# LINE BELOW PUT HERE BY ENERGI HELPER SCRIPTS (UNOFFICIAL)" >> ~/.bashrc
echo "PATH=\${PATH}:${PWD}" >> ~/.bashrc
echo "ENERGI_WALLET_ADDR=${WALLET_ADDR}" >> ~/.bashrc
echo "" >> ~/.bashrc
echo "" >> ~/.bashrc

echo ""
echo "Please run 'source ~/.bashrc' without quotes or log out and back in again (once)."

echo ""
echo "Afterwards, please try the following:"
echo ""
echo " - energi3-unlock to unlock and start staking"
echo " - energi3-status to obtain status"
echo ""
echo "More useful commands/scripts will be added in future!"
echo ""

exit 0

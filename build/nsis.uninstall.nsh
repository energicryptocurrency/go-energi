Section "Uninstall"
  # uninstall for all users
  setShellVarContext all

  # Delete (optionally) installed files
  #{{range $}}Delete $INSTDIR\{{.}}
  #{{end}}
  Delete $INSTDIR\uninstall.exe

  # Delete install directory
  rmDir $INSTDIR

  # Delete start menu launcher
  Delete "$SMPROGRAMS\${APPNAME}\Core Node.lnk"
  Delete "$SMPROGRAMS\${APPNAME}\Attach.lnk"
  Delete "$SMPROGRAMS\${APPNAME}\Testnet Core Node.lnk"
  Delete "$SMPROGRAMS\${APPNAME}\Testnet Attach.lnk"
  Delete "$SMPROGRAMS\${APPNAME}\Simnet Core Node.lnk"
  Delete "$SMPROGRAMS\${APPNAME}\Simnet Attach.lnk"
  Delete "$SMPROGRAMS\${APPNAME}\Uninstall.lnk"
  rmDir "$SMPROGRAMS\${APPNAME}"

  # Firewall - remove rules if exists
  SimpleFC::AdvRemoveRule "Energi Gen 3 incoming peers (TCP:39797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 outgoing peers (TCP:39797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 UDP discovery (UDP:39797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Testnet incoming peers (TCP:49797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Testnet outgoing peers (TCP:49797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Testnet UDP discovery (UDP:49797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Simnet incoming peers (TCP:59797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Simnet outgoing peers (TCP:59797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Simnet UDP discovery (UDP:59797)"

  # Remove IPC endpoint (https://github.com/ethereum/EIPs/issues/147)
  ${un.EnvVarUpdate} $0 "ENERGI3_SOCKET" "R" "HKLM" "\\.\pipe\energi3.ipc"

  # Remove install directory from PATH
  Push "$INSTDIR"
  Call un.RemoveFromPath

  # Cleanup registry (deletes all sub keys)
  DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}"
SectionEnd

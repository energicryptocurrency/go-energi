Name "Energi ${MAJORVERSION}.${MINORVERSION}.${BUILDVERSION}" # VERSION variables set through command line arguments
InstallDir "$InstDir"
OutFile "${OUTPUTFILE}" # set through command line arguments

# Links for "Add/Remove Programs"
!define HELPURL "https://www.energi.world/contact-us/"
!define UPDATEURL "https://docs.energi.software/en/downloads/core-node"
!define ABOUTURL "https://docs.energi.software"
!define /date NOW "%Y%m%d"

PageEx license
  LicenseData {{.License}}
PageExEnd

# Install energi3 binary
Section "Gen 3 Core Node" GETH_IDX
  SetOutPath $INSTDIR
  file {{.EnergiCoreIcon}}

  CreateDirectory "$INSTDIR\bin"
  SetOutPath "$INSTDIR\bin"
  file {{.EnergiCore}}

  # Create start menu launcher
  createDirectory "$SMPROGRAMS\${APPNAME}"
  createShortCut "$SMPROGRAMS\${APPNAME}\Core Node.lnk" "$INSTDIR\bin\energi3.exe" "" "$INSTDIR\energi-icon.ico" 0
  createShortCut "$SMPROGRAMS\${APPNAME}\Attach.lnk" "$INSTDIR\bin\energi3.exe" "attach \\.\pipe\energi3.ipc" "" ""
  createShortCut "$SMPROGRAMS\${APPNAME}\Testnet Core Node.lnk" "$INSTDIR\bin\energi3.exe" "--testnet  -ipcpath test-energi3.ipc" "$INSTDIR\energi-icon.ico" 0
  createShortCut "$SMPROGRAMS\${APPNAME}\Testnet Attach.lnk" "$INSTDIR\bin\energi3.exe" "--testnet attach \\.\pipe\test-energi3.ipc" "" ""
  createShortCut "$SMPROGRAMS\${APPNAME}\Simnet Core Node.lnk" "$INSTDIR\bin\energi3.exe" "--simnet  -ipcpath sim-energi3.ipc" "$INSTDIR\energi-icon.ico" 0
  createShortCut "$SMPROGRAMS\${APPNAME}\Simnet Attach.lnk" "$INSTDIR\bin\energi3.exe" "--simnet attach \\.\pipe\sim-energi3.ipc" "" ""
  createShortCut "$SMPROGRAMS\${APPNAME}\Uninstall.lnk" "$INSTDIR\bin\uninstall.exe" "" "" ""

  CreateShortCut "$DESKTOP\Energi Core Node.lnk" "$INSTDIR\bin\energi3.exe" "" "$INSTDIR\energi-icon.ico" 0

  # Firewall - remove rules (if exists)
  SimpleFC::AdvRemoveRule "Energi Gen 3 incoming peers (TCP:39797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 outgoing peers (TCP:39797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 UDP discovery (UDP:39797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Testnet incoming peers (TCP:49797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Testnet outgoing peers (TCP:49797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Testnet UDP discovery (UDP:49797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Simnet incoming peers (TCP:59797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Simnet outgoing peers (TCP:59797)"
  SimpleFC::AdvRemoveRule "Energi Gen 3 Simnet UDP discovery (UDP:59797)"


  # Firewall - add rules
  SimpleFC::AdvAddRule "Energi Gen 3 incoming peers (TCP:39797)" ""  6 1 1 2147483647 1 "$INSTDIR\bin\energi3.exe" "" "" "Energi" 39797 "" "" ""
  SimpleFC::AdvAddRule "Energi Gen 3 outgoing peers (TCP:39797)" ""  6 2 1 2147483647 1 "$INSTDIR\bin\energi3.exe" "" "" "Energi" "" 39797 "" ""
  SimpleFC::AdvAddRule "Energi Gen 3 UDP discovery (UDP:39797)" "" 17 2 1 2147483647 1 "$INSTDIR\bin\energi3.exe" "" "" "Energi" "" 39797 "" ""
  SimpleFC::AdvAddRule "Energi Gen 3 Testnet incoming peers (TCP:49797)" ""  6 1 1 2147483647 1 "$INSTDIR\bin\energi3.exe" "" "" "Energi" 49797 "" "" ""
  SimpleFC::AdvAddRule "Energi Gen 3 Testnet outgoing peers (TCP:49797)" ""  6 2 1 2147483647 1 "$INSTDIR\bin\energi3.exe" "" "" "Energi" "" 49797 "" ""
  SimpleFC::AdvAddRule "Energi Gen 3 Testnet UDP discovery (UDP:49797)" "" 17 2 1 2147483647 1 "$INSTDIR\bin\energi3.exe" "" "" "Energi" "" 49797 "" ""
  SimpleFC::AdvAddRule "Energi Gen 3 Simnet incoming peers (TCP:59797)" ""  6 1 1 2147483647 1 "$INSTDIR\bin\energi3.exe" "" "" "Energi" 59797 "" "" ""
  SimpleFC::AdvAddRule "Energi Gen 3 Simnet outgoing peers (TCP:59797)" ""  6 2 1 2147483647 1 "$INSTDIR\bin\energi3.exe" "" "" "Energi" "" 59797 "" ""
  SimpleFC::AdvAddRule "Energi Gen 3 Simnet UDP discovery (UDP:59797)" "" 17 2 1 2147483647 1 "$INSTDIR\bin\energi3.exe" "" "" "Energi" "" 59797 "" ""

  # Set default IPC endpoint (https://github.com/ethereum/EIPs/issues/147)
  ${EnvVarUpdate} $0 "ENERGI3_SOCKET" "R" "HKLM" "\\.\pipe\energi3.ipc"
  ${EnvVarUpdate} $0 "ENERGI3_SOCKET" "A" "HKLM" "\\.\pipe\energi3.ipc"
  ${EnvVarUpdate} $0 "ENERGI3_TESTNET_SOCKET" "R" "HKLM" "\\.\pipe\test-energi3.ipc"
  ${EnvVarUpdate} $0 "ENERGI3_TESTNET_SOCKET" "A" "HKLM" "\\.\pipe\test-energi3.ipc"
  ${EnvVarUpdate} $0 "ENERGI3_SIMNET_SOCKET" "R" "HKLM" "\\.\pipe\sim-energi3.ipc"
  ${EnvVarUpdate} $0 "ENERGI3_SIMNET_SOCKET" "A" "HKLM" "\\.\pipe\sim-energi3.ipc"

  # Add instdir to PATH
  Push "$INSTDIR"
  Call AddToPath
SectionEnd

# Install optional develop tools.
#Section /o "Development tools" DEV_TOOLS_IDX
#  SetOutPath $INSTDIR
#  {{range .DevTools}}file {{.}}
#  {{end}}
#SectionEnd

# Return on top of stack the total size (as DWORD) of the selected/installed sections.
Var GetInstalledSize.total
Function GetInstalledSize
  StrCpy $GetInstalledSize.total 0

  ${if} ${SectionIsSelected} ${GETH_IDX}
    SectionGetSize ${GETH_IDX} $0
    IntOp $GetInstalledSize.total $GetInstalledSize.total + $0
  ${endif}

  #${if} ${SectionIsSelected} ${DEV_TOOLS_IDX}
  #  SectionGetSize ${DEV_TOOLS_IDX} $0
  #  IntOp $GetInstalledSize.total $GetInstalledSize.total + $0
  #${endif}

  IntFmt $GetInstalledSize.total "0x%08X" $GetInstalledSize.total
  Push $GetInstalledSize.total
FunctionEnd

# Write registry, Windows uses these values in various tools such as add/remove program.
# PowerShell: Get-ItemProperty HKLM:\Software\Wow6432Node\Microsoft\Windows\CurrentVersion\Uninstall\* | Select-Object DisplayName, InstallLocation, InstallDate | Format-Table –AutoSize
function .onInstSuccess
  # Save information in registry in HKEY_LOCAL_MACHINE branch, Windows add/remove functionality depends on this
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "DisplayName" "${GROUPNAME} - ${APPNAME} - ${DESCRIPTION}"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "UninstallString" "$\"$INSTDIR\bin\uninstall.exe$\""
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "QuietUninstallString" "$\"$INSTDIR\bin\uninstall.exe$\" /S"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "InstallLocation" "$INSTDIR"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "InstallDate" "${NOW}"
  # Wait for Alex
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "DisplayIcon" "$\"$INSTDIR\energi-icon.ico$\""
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "Publisher" "${GROUPNAME}"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "HelpLink" "${HELPURL}"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "URLUpdateInfo" "${UPDATEURL}"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "URLInfoAbout" "${ABOUTURL}"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "DisplayVersion" "${MAJORVERSION}.${MINORVERSION}.${BUILDVERSION}"
  WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "VersionMajor" ${MAJORVERSION}
  WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "VersionMinor" ${MINORVERSION}
  # There is no option for modifying or repairing the install
  WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "NoModify" 1
  WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "NoRepair" 1

  Call GetInstalledSize
  Pop $0
  WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}" "EstimatedSize" "$0"

  # Create uninstaller
  writeUninstaller "$INSTDIR\bin\uninstall.exe"
functionEnd

Page components
Page directory
Page instfiles

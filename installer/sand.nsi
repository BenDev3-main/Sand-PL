!define APP_NAME "Sand Programming Language"
!define VERSION "1.0"
!define EXENAME "sand.exe"

Name "${APP_NAME}"
OutFile "Sand_Setup_v${VERSION}.exe"

; Corregimos la ruta a Descargas usando la variable de perfil de usuario
InstallDir "$PROFILE\Downloads\SandLang"

RequestExecutionLevel user 

Page directory
Page instfiles

Section "Instalar Sand"
    SetOutPath "$INSTDIR"
    
    ; USAMOS /nonfatal por si acaso, pero verificá la ruta.
    ; Si el script está en 'installer/', y sand.exe en la raíz, usa:
    File "..\sand.exe"
    
    WriteUninstaller "$INSTDIR\uninstall.exe"
    
    DetailPrint "Sand instalado en: $INSTDIR"
SectionEnd

Section "Uninstall"
    Delete "$INSTDIR\sand.exe"
    Delete "$INSTDIR\uninstall.exe"
    RMDir "$INSTDIR"
SectionEnd
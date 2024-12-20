
rd /s /q cmd\adapter-service\dist
cd web || exit
call npm install 
call npm run build
move  dist ..\cmd\adapter-service
cd  ..

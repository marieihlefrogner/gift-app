cd ../frontend

rm -rf build
rm -rf ../backend/public

npm run build

cp -r ./build ../backend/public

cd -
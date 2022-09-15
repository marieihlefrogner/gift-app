cd ../frontend

rm -rf build
rm -rf ../backend/frontend-build

npm run build

cp -r ./build ../backend/frontend-build

cd -
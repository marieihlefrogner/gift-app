cd ../frontend

rm -rf build
npm run build

cp -r ./build ../backend/frontend-build

cd -
# victorindo-be-assignment
Buatkanlah aplikasi backend untuk :
1)	Management User.
Berupa CRUD user serta Autentikasi dan otorisasinya. Dapat menggunakan jwt atau oauth2.
2)	Management Business Partners.
Berupa CRUD, serta Business Partners terdapat 3 jenis yakni : Customer, Supplier dan Affiliate
3)	Management Produk.
Berupa CRUD, serta Produk terdapat 3 jenis yakni, item Jual, Item Assembly dan Item Asset.

Serta tentukan infrastruktur yang sesuai untuk membangun aplikasi ini yang kemudian harinya dapat diskalakan. 
Nb: Bahasa pemograman dapat digunakan sesuai dengan yang dikuasai.

Program yang dibuat menggunakan Clean Architecture, Clean Architecture sangat disarankan untuk pengembangan aplikasi berskala besar setiap modul tersusun lebih rapi sehingga memudahkan dalam melakukan debug error, melakukan pengujian (testing) dan struktur file yang rapi juga memudahkan dalam proses pengembangan aplikasi. Sehingga Infrastruktur yang paling cocok ialah CI/CD pipeline Untuk otomatisasi build, testing, dan deployment atau dapat juga menggunakan Containerization dan jika applikasi semakin banyak diakses oleh user dapat menggunakan load balancer agar aplikasi tetap stabil ketika diakses oleh banyak user.

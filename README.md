##Scorp nedir ?
Malum Scorp un api si yok ve count sayılarını kendimiz almak zorunda kalıyoruz bu sistem size scorptan kullanıcı adı ve şifrenizi girerek istediğiniz hesapların bilgilerini çekiyor.


##How to install:

Go ile calistigi icin bilgisayarinizda go yuklu olmak zorunda

Go yu yukledikten sonra aşadaki kodu çalıştırın.

```bash
go get github.com/Sinojin/Scorp

```


##Kullanimi 

Apiyi asadaki sekilde baslatiyoruz 

```bash
	app := scorp.Api{}
	app = app.Start("KullaniciAdi","Sifre")
```

Burda onemli olan userlarin bilgilerine id leri sayesinde ulasiyoruz bunuda scropun sitesindeki profilene girdiginde url de yazili olan id yi girdigimizde hersey islem tamamlaniyor


```bash
	account := app.GetUser(1+i)
	fmt.Printf("Kullanici ID : %v Kullanici adi : %v  Isim Soyisim : %v %v  Begeni : %v ",account.User.Id,account.User.Username,account.User.FirstName,account.User.LastName,account.like_count)
	```
	
	
Bu kod ise requestler bittiginde sistemden cikis yapmayi sagliyor koymasanizda olur ama oneririm sistemden cikis yapiniz.
	
	
	```bash
	app.Close()
	```

##Ornek Kullanim

Burda Yapmak istedigim 0 ile 10 arasindaki  idlere  sahip herkesi terminale basmak


 	```bash	
 	
	app := scorp.Api{}
	app = app.Start("kullaniciAdi","Sifre")

    for  i:=0;i<10 ;i++  {
		account := app.GetUser(1+i)
		fmt.Printf("Kullanici ID : %v Kullanici adi : %v  Isim Soyisim : %v %v  Begeni : %v ",account.User.Id,account.User.Username,account.User.FirstName,account.User.LastName,account.like_count)
	
	}

	println()
	go app.Close()

	```
	
Burdan Sonrasi sizin hayal gucunuze kalmis ilerleyen zamanlarda yorumlari ve takipcileri cekmekde yapabilirim tabiki istek olursa herkese kolay gelsin.





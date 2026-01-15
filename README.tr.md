# ZeroGate Tekil İşlem Ödeme Geçidi API'si

[Click here for English Document (İngilizce Belge İçin Tıklayın)](README.md)

ZeroGate, finansal işlemlerde çift çekimi (double-spending) önlemek için Redis tabanlı dağıtık kilitler ve idempotency (tekillik) deseni kullanan bir ödeme geçidi bileşenidir.

## Mimari
Sistem, kalıcı veri saklama için PostgreSQL ve işlemlerden önce dağıtık kilit almak için Redis kullanır. Bu sayede aynı anda gelen kopya istekler sıraya alınır veya reddedilir, böylece veri bütünlüğü sağlanır.

## Nasıl Kurulur ve Çalıştırılır?

Projeyi lokal bilgisayarınızda çalıştırmak için aşağıdaki adımları izleyin:

1. Veritabanı ve Önbellek Sunucularını Başlatın:
   Docker kullanarak PostgreSQL ve Redis'i arka planda çalıştırın.
   ```bash
   docker-compose up -d
   ```

2. Gerekli Kütüphaneleri İndirin:
   ```bash
   go mod tidy
   ```

3. Uygulamayı Başlatın:
   ```bash
   go run cmd/api/main.go
   ```
   Uygulama başarıyla başladığında terminalde `Server listening on port 8080` mesajını göreceksiniz.

## Nasıl Test Edilir? (Idempotency Kanıtı)

Uygulama çalışırken YENİ bir terminal penceresi açın ve aşağıdaki adımları sırasıyla uygulayın.

1. **İlk İstek (Normal Ödeme İşlemi):**
   Aşağıdaki komutu çalıştırarak bir ödeme isteği gönderin.
   ```bash
   curl -X POST http://localhost:8080/api/pay \
        -H "Content-Type: application/json" \
        -d '{"idempotency_key": "MUSTERI-ODEME-001", "amount": 250.00}'
   ```
   *Beklenen Sonuç:* İşlem onaylanacak, yeni bir ID üretilecek ve `status: completed` dönecektir.

2. **İkinci İstek (Çift Çekim Engelleme Testi):**
   Aynı komutu BİREBİR AYNI ŞEKİLDE tekrar çalıştırın (kullanıcının öde butonuna ikinci kez basmasını simüle ediyoruz).
   ```bash
   curl -X POST http://localhost:8080/api/pay \
        -H "Content-Type: application/json" \
        -d '{"idempotency_key": "MUSTERI-ODEME-001", "amount": 250.00}'
   ```
   *Beklenen Sonuç:* Sistem parayı ikinci kez çekmeyecektir. İlk işlemde üretilen ID'nin BİREBİR AYNISINI ekrana basarak işlemin zaten gerçekleştiğini belirtecektir.

## Uç Noktalar
POST /api/pay
Idempotency anahtarı ve miktar kabul eder. Her anahtar için sadece bir işlemin gerçekleşeceğini garanti eder.

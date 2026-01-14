# ZeroGate Tekil İşlem Ödeme Geçidi API'si

[Click here for English Document (İngilizce Belge İçin Tıklayın)](README.md)

ZeroGate, finansal işlemlerde çift çekimi (double-spending) önlemek için Redis tabanlı dağıtık kilitler ve idempotency (tekillik) deseni kullanan bir ödeme geçidi bileşenidir.

## Mimari
Sistem, kalıcı veri saklama için PostgreSQL ve işlemlerden önce dağıtık kilit almak için Redis kullanır. Bu sayede aynı anda gelen kopya istekler sıraya alınır veya reddedilir, böylece veri bütünlüğü sağlanır.

## Uç Noktalar
POST /api/pay
Idempotency anahtarı ve miktar kabul eder. Her anahtar için sadece bir işlemin gerçekleşeceğini garanti eder.

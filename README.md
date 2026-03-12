SSH Connections TUI
Aplikacja terminalowa napisana w Go, umożliwiająca wygodne zarządzanie hostami SSH z poziomu interfejsu tekstowego (TUI). Pozwala szybko organizować serwery, przeglądać ich szczegóły oraz łączyć się jednym klawiszem.

✨ Funkcje
grupowanie hostów w sekcje

podgląd szczegółów hosta

szybkie uruchamianie połączeń SSH

wykonywanie ping bez opuszczania aplikacji

interfejs oparty na bibliotekach tview i tcell

🧩 Wymagania
Przed uruchomieniem upewnij się, że masz zainstalowane:

Go 1.20+

ssh

ping

Sprawdzenie wersji:

go version
ssh -V
ping -V

📦 Instalacja
Sklonuj repozytorium:

git clone https://github.com/twoj-user/sshcm.git
cd sshcm

Pobierz zależności:

go mod tidy

Uruchomienie bez budowania:

go run main.go

Zbuduj program:

go build -o sshcm

🚀 Uruchomienie
Program wymaga pliku konfiguracyjnego:

ssh_connections.conf
mv ssh_connections.conf.example ssh_connections.conf

Uruchomienie aplikacji:

./sshcm

⚙️ Konfiguracja
Aplikacja korzysta z pliku:

ssh_connections.conf

Przykładowa konfiguracja
group Production Servers

host web1
hostname 192.168.1.10
user root
port 22
description Główny serwer WWW

host db1
hostname 192.168.1.20
user admin
port 22
description Serwer bazy danych

group Development

host dev1
hostname 192.168.1.30
user dev
description Serwer developerski

🎮 Sterowanie
Klawisz / Akcja
↑ / ↓ – nawigacja
→ – rozwiń grupę
← – zwiń grupę
Enter – połącz przez SSH
p – ping hosta

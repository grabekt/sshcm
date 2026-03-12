# SSH Connections manager (sshcm)

Aplikacja terminalowa napisana w **Go**, umożliwiająca wygodne zarządzanie hostami SSH z poziomu interfejsu tekstowego (TUI). Pozwala szybko organizować serwery, przeglądać ich szczegóły oraz łączyć się z nimi jednym klawiszem.

## ✨ Funkcje

* **Grupowanie hostów:** Organizacja serwerów w sekcje (np. Produkcja, Dev).
* **Podgląd szczegółów:** Szybki wgląd w parametry wybranego hosta.
* **Szybkie połączenia:** Uruchamianie sesji SSH bezpośrednio z aplikacji.
* **Wbudowany Ping:** Sprawdzanie dostępności hosta bez zamykania TUI.
* **Lekki interfejs:** Oparty na sprawdzonych bibliotekach `tview` oraz `tcell`.

## 🧩 Wymagania

Upewnij się, że w Twoim systemie znajdują się:
   * **Go 1.20+**
   * **ssh**
   * **ping**

Sprawdzenie wersji:

   go version
   ssh -V
   ping -V

## 📦 Instalacja

1. Zainstaluj golang:
   
   sudo apt update
   
   sudo apt install golang-go

2. Pobrac repozytorium
   
   git clone adresrepo
   
   cd sshcm
   
3. Pobierz zależności:
   
   go mod tidy
   
4. Zbuduj program:
   
   go build -o sshcm

5. Uruchomienie bez budowania:

   go run main.go
   
   
## 🚀 Uruchomienie

1. Przygotuj plik konfiguracyjny:
   
   cp ssh_connections.conf.example ssh_connections.conf
   
2. Uruchom aplikację:
   
   ./sshcm
   
   
## ⚙️ Konfiguracja
Aplikacja korzysta z pliku ssh_connections.conf. Przykładowa struktura:

```
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
```

## 🎮 Sterowanie

| Klawisz | Akcja |
|---|---|
| ↑ / ↓ | Nawigacja po liście |
| → | Rozwinięcie grupy |
| ← | Zwinięcie grupy |
| Enter | Połącz przez SSH |
| p | Wykonaj Ping na hoście |
| q / Ctrl+C | Wyjście z aplikacji |

------------------------------


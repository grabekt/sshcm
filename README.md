SSH Connections TUI

Terminalowa aplikacja napisana w Go, umożliwiająca zarządzanie i łączenie się z hostami SSH z poziomu wygodnego interfejsu tekstowego (TUI).

Program pozwala:

organizować hosty w grupy

przeglądać szczegóły hosta

uruchamiać SSH jednym klawiszem

wykonywać ping do hosta

Interfejs jest zbudowany przy użyciu bibliotek:

tview

tcell

Wymagania

Przed uruchomieniem upewnij się, że masz zainstalowane:

Go 1.20+

ssh

ping

Sprawdzenie:

go version
ssh -V
ping -V
Instalacja

Sklonuj repozytorium:

git clone https://github.com/twoj-user/ssh-connections-tui.git
cd ssh-connections-tui

Pobierz zależności:

go mod tidy

Zbuduj program:

go build -o ssh-tui
Uruchomienie

Program wymaga pliku konfiguracyjnego:

ssh_connections.conf

Uruchomienie:

./ssh-tui
Konfiguracja

Program korzysta z pliku:

ssh_connections.conf
Przykład konfiguracji
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
Sterowanie
Klawisz	Akcja
↑ / ↓	nawigacja
→	rozwiń grupę
←	zwiń grupę
Enter	połącz przez SSH
p	ping hosta

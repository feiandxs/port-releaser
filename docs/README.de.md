# Port Releaser

Ein einfaches Tool zum Freigeben belegter Ports.

## Installation

```bash
go install github.com/codenamev/port-releaser@latest
```

## Verwendung

```bash
port-releaser -p 3000
```

## Parameter

- `-p`: Die Portnummer, die freigegeben werden soll

## Funktionsweise

1. Findet den Prozess, der den angegebenen Port verwendet, mit dem `netstat`-Befehl
2. Beendet den Prozess mit dem `taskkill`-Befehl
3. Wenn der Prozess nicht normal beendet werden kann, wird der Parameter `-f` verwendet, um die Beendigung zu erzwingen

## Vorsichtsmaßnahmen

- Administratorrechte sind erforderlich, um bestimmte Prozesse zu beenden
- Mit Vorsicht verwenden und sicherstellen, dass der zu beendende Prozess kein kritischer Systemprozess ist

## Lizenz

Dieses Projekt steht unter der Creative Commons Namensnennung-Nicht kommerziell 4.0 International Lizenz (CC BY-NC 4.0).

Sie dürfen:
- Teilen — das Material in jedwedem Format oder Medium vervielfältigen und weiterverbreiten
- Bearbeiten — das Material remixen, verändern und darauf aufbauen

Unter folgenden Bedingungen:
- Namensnennung — Sie müssen angemessene Urheber- und Rechteangaben machen, einen Link zur Lizenz beifügen und angeben, ob Änderungen vorgenommen wurden
- Nicht kommerziell — Sie dürfen das Material nicht für kommerzielle Zwecke nutzen ohne ausdrückliche Erlaubnis des Autors

Weitere Informationen: https://creativecommons.org/licenses/by-nc/4.0/deed.de

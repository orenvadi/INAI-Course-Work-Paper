#import "@preview/codelst:2.0.0": sourcecode, sourcefile


#heading(
level: 1, outlined: true,
    numbering: (..nums) => "Kapitel " + nums
      .pos()
      .map(str)
      .join(".") + "."
  )[Server]


Zur Speicherung und Verarbeitung von Daten wurde ein Server erstellt


== Serverfunktionen

- Speicherung aller Daten
- Bereitstellung des Zugriffs auf Daten




== Projektarchitektur
Datenbank: SurrealDB @surrealdb-doc

Programmiersprache: Go @golang-doc


=== Projektstruktur

#sourcefile(read("../images/servertree.sh"), file:"../images/servertree.sh")
#align(center)[Listing 5: Server-Projektstruktur]


== Beschreibung des Quellcodes

=== main.go


#sourcefile(read("../../GoBackend/cmd/sso/main.go"), file:"../../GoBackend/cmd/sso/main.go")
#align(center)[Listing 6: Databank Schema]

*1. Importe:*

- Der Code beginnt mit dem Importieren der erforderlichen Pakete:
- `fmt`: für formatiertes Drucken
- `log/slog`: für strukturiertes Protokollieren
- `os`: für die Interaktion mit dem Betriebssystem
- `os/signal`: für die Verarbeitung von Signalen vom Betriebssystem
- `syscall`: für Systemaufrufe (wie Signalverarbeitung)
- Zwei interne Pakete (`app` und `config`), die für diese Anwendung spezifisch sind

*2. Konfiguration:*

- Die Funktion `main` beginnt mit dem Laden der Konfiguration mit `config.MustLoad()`. Diese Funktion liest wahrscheinlich Konfigurationseinstellungen aus einer Umgebungsdatei oder einer anderen Quelle und gibt ein Konfigurationsobjekt zurück.
- Die geladene Konfiguration wird dann mit `fmt.Printf` gedruckt.

*3. Logger-Setup:*

- Eine Funktion namens „setupLogger“ wird verwendet, um den Logger basierend auf der Umgebung (`cfg.Env`) zu konfigurieren.

- Die Funktion verwendet eine Switch-Anweisung, um je nach Umgebung den geeigneten Protokollierungshandler auszuwählen:

- Im lokalen Modus (`envLocal`) wird ein Texthandler verwendet, der alles (Debug-Level) in die Standardausgabe (`os.Stdout`) protokolliert.

- Im Entwicklungsmodus (`envDev`) wird ein JSON-Handler verwendet, der ebenfalls alles (Debug-Level) in die Standardausgabe protokolliert.

- Im Produktionsmodus (`envProd`) wird ein JSON-Handler verwendet, der nur Informationsmeldungen und höher in die Standardausgabe protokolliert.

*4. Anwendungsinitialisierung:*

- Der Code erstellt eine neue Anwendungsinstanz mit „app.New“. Diese Funktion verwendet wahrscheinlich den Logger, den gRPC-Port, das Konfigurationsobjekt und die Token-TTL als Argumente und gibt ein Anwendungsobjekt zurück.

*5. Ausführen des gRPC-Servers:*

- Der Code startet den gRPC-Server der Anwendung in einer Goroutine mit `application.GRPCSrv.MustRun()`. Diese Funktion startet wahrscheinlich den gRPC-Server auf dem konfigurierten Port und bearbeitet gleichzeitig eingehende Anfragen.

*6. Ordentliches Herunterfahren:*

- Ein Kanal `stop` vom Typ `os.Signal` wird erstellt, um Signale vom Betriebssystem zu empfangen.

- Die Funktion `signal.Notify` wird verwendet, um Beendigungssignale (`SIGINT` und `SIGTERM`) zu registrieren und sie an den `stop`-Kanal zu senden.

- Der Code blockiert dann das Lesen vom `stop`-Kanal mit `<-stop`. Dies wartet, bis ein Signal empfangen wird.

- Sobald ein Signal empfangen wird, protokolliert der Code Informationen über das Signal mit dem konfigurierten Logger.

- Die Methode `application.GRPCSrv.Stop()` wird aufgerufen, um den gRPC-Server ordnungsgemäß zu stoppen.

- Schließlich protokolliert der Code eine Meldung, die angibt, dass die Anwendung gestoppt wurde.


=== Schema.sql


#sourcefile(read("../../GoBackend/migrations/surreal/1_init.up.surql"), file:"../../GoBackend/migrations/surreal/1_init.up.surql")
#align(center)[Listing 7: Databank Schema]



Dieser Code definiert das Schema für mehrere Tabellen, die zur Verwaltung einer Schul- oder Universitätsumgebung verwendet werden könnten. Hier ist eine Aufschlüsselung der einzelnen Tabellen und ihrer Felder:

*Tabellen:*

- *Student:* Speichert Informationen über Studenten.
- *Name:* Flexibles Objekt zur Speicherung detaillierter Studenteninformationen (wahrscheinlich Name, Adresse usw.).
- *StudentCode:* Eindeutige Kennung für den Studenten (wahrscheinlich eine Zeichenfolge-ID). (Indiziert für schnellere Nachschlagevorgänge)
- *Email:* E-Mail-Adresse des Studenten. (Indiziert für schnellere Nachschlagevorgänge)
- *PasswordHash:* Hash-Passwort für sichere Authentifizierung.
- *Fächer:* Array von Datensätzen mit Details zu den Fächern, in die ein Student eingeschrieben ist. (verweist wahrscheinlich auf die Fachtabelle)
- *Gruppen:* Array von Datensätzen mit Details zu den Gruppen, zu denen ein Student gehört. (verweist wahrscheinlich auf die Tabelle „Gruppe“)
- *Lehrer:* Speichert Informationen über Lehrer.
- *Name:* Flexibles Objekt zur Speicherung detaillierter Lehrerinformationen.
- *Lehrercode:* Eindeutige Kennung für den Lehrer (wahrscheinlich eine Zeichenfolge-ID). (Indiziert für schnellere Nachschlagevorgänge)
- *E-Mail:* E-Mail-Adresse des Lehrers. (Indiziert für schnellere Nachschlagevorgänge)
- *PasswortHash:* Hash-Passwort für sichere Authentifizierung.
- *Fächer:* Array von Datensätzen mit Details zu den Fächern, die ein Lehrer unterrichtet. (verweist wahrscheinlich auf die Tabelle „Fächer“)
- *Gruppen:* Array von Datensätzen mit Details zu den Gruppen, die ein Lehrer verwaltet. (verweist wahrscheinlich auf die Tabelle „Gruppe“)
- *Fach:* Speichert Informationen zu angebotenen Fächern.
- *Name:* Name des Fachs.
- *Gruppe:* Speichert Informationen zu Gruppen innerhalb der Schule.
- *Name:* Name der Gruppe. (Indiziert für schnellere Nachschlagevorgänge)
- *Standort:* Speichert Informationen zu Orten, an denen der Unterricht stattfindet.
- *Name:* Name des Standorts (z. B. Name des Klassenzimmers). (Indiziert für schnellere Nachschlagevorgänge)
- *PrelimConfirmCode:* Optionaler Code für vorläufige Bestätigung (Zweck nicht ganz klar aus dem Schema).
- *Stundenplan:* Speichert Informationen zu Stundenplänen.



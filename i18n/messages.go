package i18n

// Messages 存储所有语言的消息
var Messages = map[string]map[string]string{
	"en": {
		"specifyPort":       "Please specify a port number.",
		"unsupportedOS":     "Unsupported operating system: %s\n",
		"errorFindingPort":  "Error finding port: %v\n",
		"noProcessFound":    "No process found occupying port %s.\n",
		"process":           "Process: %s\n",
		"terminateProcess":  "Do you want to terminate this process? [y/N]: ",
		"unableToTerminate": "Unable to terminate process %s: %v\n",
		"processTerminated": "Process %s has been terminated.\n",
	},
	"zh": {
		"specifyPort":       "请指定一个端口号。",
		"unsupportedOS":     "不支持的操作系统: %s\n",
		"errorFindingPort":  "查找端口时发生错误: %v\n",
		"noProcessFound":    "没有找到占用端口 %s 的进程。\n",
		"process":           "进程: %s\n",
		"terminateProcess":  "是否终止此进程? [y/N]: ",
		"unableToTerminate": "无法终止进程 %s: %v\n",
		"processTerminated": "进程 %s 已终止。\n",
	},
	"ja": {
		"specifyPort":       "ポート番号を指定してください。",
		"unsupportedOS":     "サポートされていないオペレーティングシステム: %s\n",
		"errorFindingPort":  "ポートの検索中にエラーが発生しました: %v\n",
		"noProcessFound":    "ポート %s を使用しているプロセスが見つかりません。\n",
		"process":           "プロセス: %s\n",
		"terminateProcess":  "このプロセスを終了しますか? [y/N]: ",
		"unableToTerminate": "プロセス %s を終了できません: %v\n",
		"processTerminated": "プロセス %s が終了しました。\n",
	},
	"es": {
		"specifyPort":       "Por favor, especifique un número de puerto.",
		"unsupportedOS":     "Sistema operativo no soportado: %s\n",
		"errorFindingPort":  "Error al buscar el puerto: %v\n",
		"noProcessFound":    "No se encontró ningún proceso ocupando el puerto %s.\n",
		"process":           "Proceso: %s\n",
		"terminateProcess":  "¿Desea terminar este proceso? [y/N]: ",
		"unableToTerminate": "No se puede terminar el proceso %s: %v\n",
		"processTerminated": "El proceso %s ha sido terminado.\n",
	},
	"fr": {
		"specifyPort":       "Veuillez spécifier un numéro de port.",
		"unsupportedOS":     "Système d'exploitation non pris en charge : %s\n",
		"errorFindingPort":  "Erreur lors de la recherche du port : %v\n",
		"noProcessFound":    "Aucun processus trouvé occupant le port %s.\n",
		"process":           "Processus : %s\n",
		"terminateProcess":  "Voulez-vous terminer ce processus ? [y/N] : ",
		"unableToTerminate": "Impossible de terminer le processus %s : %v\n",
		"processTerminated": "Le processus %s a été terminé.\n",
	},
	"de": {
		"specifyPort":       "Bitte geben Sie eine Portnummer an.",
		"unsupportedOS":     "Nicht unterstütztes Betriebssystem: %s\n",
		"errorFindingPort":  "Fehler beim Suchen des Ports: %v\n",
		"noProcessFound":    "Kein Prozess gefunden, der Port %s belegt.\n",
		"process":           "Prozess: %s\n",
		"terminateProcess":  "Möchten Sie diesen Prozess beenden? [y/N]: ",
		"unableToTerminate": "Prozess %s kann nicht beendet werden: %v\n",
		"processTerminated": "Prozess %s wurde beendet.\n",
	},
}
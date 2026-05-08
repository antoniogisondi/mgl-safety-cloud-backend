package seeders

import (
	"log"

	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/database"
	"github.com/antoniogisondi/mgl-safety-cloud-backend/internal/models"
)

func SeedDeadlineCategories() {
	categories := []models.DeadlineCategory{
		{Name: "DVR", Group: "Documenti sicurezza", Description: "Documento di Valutazione dei Rischi", ValidityMonths: 0},
		{Name: "DUVRI", Group: "Documenti sicurezza", Description: "Documento Unico di Valutazione dei Rischi Interferenziali", ValidityMonths: 0},
		{Name: "POS", Group: "Cantieri", Description: "Piano Operativo di Sicurezza", ValidityMonths: 0},
		{Name: "PSC", Group: "Cantieri", Description: "Piano di Sicurezza e Coordinamento", ValidityMonths: 0},
		{Name: "PIMUS", Group: "Cantieri", Description: "Piano di Montaggio, Uso e Smontaggio ponteggi", ValidityMonths: 0},
		{Name: "Piano emergenza", Group: "Documenti sicurezza", Description: "Piano di emergenza ed evacuazione", ValidityMonths: 0},
		{Name: "Registro antincendio", Group: "Antincendio", Description: "Registro controlli antincendio", ValidityMonths: 0},

		{Name: "Formazione lavoratori", Group: "Formazione", Description: "Formazione generale e specifica lavoratori", ValidityMonths: 60},
		{Name: "Aggiornamento lavoratori", Group: "Formazione", Description: "Aggiornamento quinquennale lavoratori", ValidityMonths: 60},
		{Name: "Formazione preposti", Group: "Formazione", Description: "Formazione particolare aggiuntiva per preposti", ValidityMonths: 24},
		{Name: "Formazione dirigenti", Group: "Formazione", Description: "Formazione dirigenti sicurezza", ValidityMonths: 60},
		{Name: "RSPP Datore di lavoro", Group: "Formazione", Description: "Corso datore di lavoro RSPP", ValidityMonths: 60},
		{Name: "RLS", Group: "Formazione", Description: "Formazione e aggiornamento RLS", ValidityMonths: 12},
		{Name: "Antincendio", Group: "Formazione", Description: "Formazione addetti antincendio", ValidityMonths: 36},
		{Name: "Primo soccorso", Group: "Formazione", Description: "Formazione addetti primo soccorso", ValidityMonths: 36},
		{Name: "Carrelli elevatori", Group: "Attrezzature", Description: "Abilitazione conduzione carrelli elevatori", ValidityMonths: 60},
		{Name: "PLE", Group: "Attrezzature", Description: "Abilitazione piattaforme di lavoro elevabili", ValidityMonths: 60},
		{Name: "Trattori agricoli", Group: "Agricoltura", Description: "Abilitazione conduzione trattori agricoli o forestali", ValidityMonths: 60},
		{Name: "Ponteggi", Group: "Cantieri", Description: "Formazione addetti montaggio/smontaggio ponteggi", ValidityMonths: 48},
		{Name: "Lavori in quota", Group: "Formazione", Description: "Formazione lavori in quota e DPI III categoria", ValidityMonths: 60},
		{Name: "PES PAV PEI", Group: "Formazione", Description: "Formazione rischio elettrico", ValidityMonths: 60},
		{Name: "HACCP", Group: "Igiene alimentare", Description: "Formazione igiene alimentare", ValidityMonths: 24},

		{Name: "Visita medica preventiva", Group: "Sorveglianza sanitaria", Description: "Visita medica preventiva lavoratore", ValidityMonths: 0},
		{Name: "Visita medica periodica", Group: "Sorveglianza sanitaria", Description: "Visita medica periodica lavoratore", ValidityMonths: 12},
		{Name: "Idoneità sanitaria", Group: "Sorveglianza sanitaria", Description: "Giudizio di idoneità alla mansione", ValidityMonths: 12},

		{Name: "Estintori", Group: "Antincendio", Description: "Controllo periodico estintori", ValidityMonths: 6},
		{Name: "Impianto di terra", Group: "Impianti", Description: "Verifica periodica impianto di messa a terra", ValidityMonths: 60},
		{Name: "Impianto elettrico", Group: "Impianti", Description: "Controllo impianto elettrico", ValidityMonths: 0},
		{Name: "Porte REI", Group: "Antincendio", Description: "Controllo porte resistenti al fuoco", ValidityMonths: 6},
		{Name: "Attrezzature di lavoro", Group: "Attrezzature", Description: "Manutenzione e verifica attrezzature di lavoro", ValidityMonths: 12},

		{Name: "Consegna DPI", Group: "DPI", Description: "Verbale consegna dispositivi di protezione individuale", ValidityMonths: 0},
		{Name: "Controllo DPI III categoria", Group: "DPI", Description: "Controllo DPI anticaduta e III categoria", ValidityMonths: 12},

		{Name: "Analisi acqua", Group: "Igiene alimentare", Description: "Analisi acqua potabile", ValidityMonths: 12},
		{Name: "Derattizzazione", Group: "Igiene e ambiente", Description: "Interventi di derattizzazione e monitoraggio infestanti", ValidityMonths: 6},
		{Name: "Sanificazione", Group: "Igiene e ambiente", Description: "Interventi di sanificazione ambienti", ValidityMonths: 0},

		{Name: "Revisione trattori", Group: "Agricoltura", Description: "Revisione periodica trattori agricoli", ValidityMonths: 0},
		{Name: "Fitosanitari", Group: "Agricoltura", Description: "Abilitazione acquisto e utilizzo prodotti fitosanitari", ValidityMonths: 60},
		{Name: "Controllo atomizzatori", Group: "Agricoltura", Description: "Controllo funzionale macchine irroratrici", ValidityMonths: 36},

		{Name: "Piano safety", Group: "Eventi", Description: "Piano safety per eventi e manifestazioni", ValidityMonths: 0},
		{Name: "Piano security", Group: "Eventi", Description: "Piano security per eventi e manifestazioni", ValidityMonths: 0},
		{Name: "Collaudi strutture", Group: "Eventi", Description: "Collaudi strutture temporanee per eventi", ValidityMonths: 0},
	}

	for _, category := range categories {
		var existing models.DeadlineCategory

		err := database.DB.Where("name = ?", category.Name).First(&existing).Error

		if err != nil {
			if err := database.DB.Create(&category).Error; err != nil {
				log.Println("Errore creazione categoria:", category.Name)
			}
		}
	}

	log.Println("Seeder categorie scadenze completato")
}

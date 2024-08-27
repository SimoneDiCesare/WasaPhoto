<template>
  <div class="error-container">
    <h1>{{ errorMessage }}</h1>
  </div>
</template>

<script>
export default {
  computed: {
    errorMessage() {
      // Leggi il codice dalla query string
      const code = this.$route.query.code;

      // Definisci i messaggi di errore in base al codice
      const errorMessages = {
        '404': 'Pagina non trovata',
        '500': 'Errore interno del server',
        '403': 'Accesso negato',
        // Aggiungi altri codici e messaggi di errore qui
        default: 'Si è verificato un errore'
      };

      // Ritorna il messaggio corrispondente al codice, oppure un messaggio di default
      return errorMessages[code] || errorMessages['default'];
    }
  }
}
</script>

<style scoped>
.error-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  text-align: center;
}
</style>

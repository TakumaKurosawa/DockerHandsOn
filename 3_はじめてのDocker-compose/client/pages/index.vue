<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm8 md6>
      <div class="text-center">
        <logo />
        <vuetify-logo />
      </div>
      <v-card>
        <v-card-title class="headline">
          サーバサイドレンダリングでのAPIレスポンス
        </v-card-title>
        <v-card-text>
          <h2>
            {{ message }}
          </h2>
        </v-card-text>
      </v-card>
      <v-card>
        <v-card-title class="headline">
          ブラウザからリクエストを投げた結果
        </v-card-title>
        <v-card-text>
          <h2>
            {{ response ? response : 'none' }}
          </h2>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="primary" @click="getMessage">
            Continue
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'

export default {
  components: {
    Logo,
    VuetifyLogo,
  },
  async asyncData({ $axios }) {
    return await $axios
      .$get('http://api:8080/server/access')
      .catch((err) => err)
  },
  data() {
    return {
      response: '',
    }
  },
  methods: {
    async getMessage() {
      const { message } = await this.$axios
        .$get('http://localhost:8080/browser/access')
        .catch((err) => err)
      this.response = message
    },
  },
}
</script>

<template>
  <v-container>
    <v-row>
      <v-col cols="12" md="4">
        <v-form ref="form" lazy-validation>
          <v-btn class="mr-4" @click="submit">
            <span>戻る</span>
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Result",
  data() {
    return {};
  },
  methods: {
    submit() {
      this.$router.push({
        name: "Main",
      });
    },
  },
};
</script>

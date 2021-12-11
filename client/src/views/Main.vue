<template>
  <v-container>
    <v-row>
      <v-col cols="12" md="4">
        <v-form ref="form" lazy-validation>
          <v-text-field
            v-model="inputName"
            :counter="20"
            label="名前を入れてね"
            placeholder="例：おみくじ 太郎"
            required
            outlined
          ></v-text-field>
          <v-btn class="mr-4" :disabled="isValid" @click="submitForm">
            <span>占う！</span>
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      name: "",
      success: true,
      isValid: false,
      // 入力規則
      //name_required: (value) => !!value || "名前が無いと占えないよ", // 入力必須の制約
      //limit_length: (value) => value.length <= 20 || "20文字以内で入力してね", // 文字数の制約
    };
  },
  methods: {
    submitForm() {
      //if (this.$refs.form.validate()) {
      console.log("OK");
      this.$router.push({
        name: "Result",
      });
      // } else {
      //   console.log("NG");
      // }
    },
  },
};
</script>

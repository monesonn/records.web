Vue.use(VueMaterial.default);
Vue.use(window.vuelidate.default);
// Vue.use(window.vuelidate.validators);

const { required, email, maxLength, minLength } = window.validators;

Vue.component("login-comp", {
  template: `<div accept-charset="utf-8">
                <form novalidate @submit.prevent="validateLogin">
                    <div class="group">
                        <md-field :class="getValidation('email')">
                            <label>Email</label>
                            <md-input v-model="loginData.email" type="email"></md-input>
                            <span class="md-error" v-if="!$v.loginData.email.required">Enter email</span>
                            <span class="md-error" v-if="!$v.loginData.email.email">Not valid email</span>
                        </md-field>
                    </div>
                    <div class="group">
                        <md-field :class="getValidation('password')">
                            <label>Password</label>
                            <md-input v-model="loginData.password" type="password"></md-input>
                            <span class="md-error" v-if="!$v.loginData.password.required">Enter password</span>
                        </md-field>
                    </div>
                <md-progress-bar md-mode="indeterminate" v-if="sending" />
                <div v-if="err === null">
                <button type="submit" class="btn btn-outline-primary is-outlined" style="width:200px;" :disabled="sending">Увійти</button>
                </button>
                </div>
                <div v-if="err != null" style="color:red; padding: 10px 0;">
                    <button type="submit" class="btn btn-outline-danger is-outlined" style="margin: 15px 0;" :disabled="sending">Увійти</button>
                    <small>{{ err.msg }}</small>
                </div>
                </form>
            </div>`,
  data: function () {
    return {
      loginData: {
        email: "",
        password: "",
      },
      err: null,
      sending: false,
    };
  },
  validations: {
    loginData: {
      email: { required, email },
      password: { required },
    },
  },
  methods: {
    getValidation(fieldName) {
      const field = this.$v.loginData[fieldName];

      if (field) {
        return {
          "md-invalid": field.$invalid && field.$dirty,
        };
      }
    },
    clearForm: function () {
      this.$v.$reset();
      this.loginData.email = "";
      this.loginData.password = "";
    },
    login: function () {
      this.sending = true;
      axios
        .post("/api/sign/in", this.loginData)
        .then((res) => {
          if (res.data.admin === true) {
            localStorage.access = res.data.tokens.access;
            localStorage.refresh = res.data.tokens.refresh;
            app.isLogin = true;
            console.log("Success");
          } else {
            this.err = "You are not administrator.";
          }
          this.sending = false;
        })
        .catch((err) => {
          this.err = err.response.data;
          this.sending = false;
        });
    },
    validateLogin() {
      this.$v.$touch();

      if (!this.$v.$invalid) {
        this.login();
      }
    },
  },
  watch: {
    email(email) {
      this.loginData.email = email;
    },
    password(password) {
      this.loginData.password = password;
    },
  },
});

var app = new Vue({
  el: "#app",
  data: {
    data: "Hello",
    isLogin: true,
    menuOpen: false,
  },
  methods: {
    init: function () {},
  },
});

app.init();

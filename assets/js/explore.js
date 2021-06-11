Vue.use(VueMaterial.default);
Vue.use(window.vuelidate.default);
// Vue.use(window.vuelidate.validators);

const { required, email, maxLength, minLength } = window.validators;

Vue.component("login-comp", {
  template: `<div accept-charset="utf-8" style="width:200px;">
                <nav>
                    <ol class="breadcrumb-list">
                        <li class="breadcrumb-item active">Вхід</li>
                        <li class="breadcrumb-item" style="cursor: pointer" @click="toggle = !toggle">Реєстрація</li>
                    </ol>
                </nav>
                <br></br>
                <form novalidate @submit.prevent="validateLogin">
                    <div class="group">
                        <md-field :class="getValidation('email')">
                            <label>Електронна пошта</label>
                            <md-input v-model="loginData.email" type="email"></md-input>
                            <span class="md-error" v-if="!$v.loginData.email.required">Введіть електронну адресу</span>
                            <span class="md-error" v-if="!$v.loginData.email.email">Це не схоже на електронну адресу</span>
                        </md-field>
                    </div>
                    <div class="group">
                        <md-field :class="getValidation('password')">
                            <label>Пароль</label>
                            <md-input v-model="loginData.password" type="password"></md-input>
                            <span class="md-error" v-if="!$v.loginData.password.required">Йой, Ви забули ввести пароль</span>
                        </md-field>
                    </div>
                <md-progress-bar md-mode="indeterminate" v-if="sending" />
                <div v-if="err === null">
                <button type="submit" class="btn btn-outline-primary is-outlined" style="width:200px;" :disabled="sending">Увійти</button>
                </button>
                </div>
                <div v-else="err != null" style="color:red; padding: 10px 0;">
                    <button type="submit" class="btn btn-outline-danger is-outlined" style="width:200px; margin: 15px 0;" :disabled="sending">Увійти</button>
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
      uuid: null,
      err: null,
      toggle: true,
      sending: false,
      sucess: false,
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
    validateLogin() {
      this.$v.$touch();

      if (!this.$v.$invalid) {
        this.login();
      }
    },
    login: function () {
      this.sending = true;
      axios
        .post("/api/sign/in", this.loginData)
        .then((res) => {
          localStorage.access = res.data.tokens.access;
          localStorage.refresh = res.data.tokens.refresh;
          this.getProfile();
        })
        .catch((err) => {
          this.err = err.response.data;
          this.sending = false;
        });
    },
    getProfile: function () {
      window.setTimeout(() => {
        axios
          .get(`/api/me/email/${this.loginData.email}`, {
            headers: {
              Authorization: `Bearer ${localStorage.access}`,
            },
          })
          .then((res) => {
            app.profile = res.data.client;
            app.profile.initials = app.profile.fname[0] + app.profile.lname[0];
            app.profile.date = app.profile.birthday.substring(0, 10);
            app.loginStatus = true;
            localStorage.uuid = app.profile.id;
            this.err = null;
          })
          .catch((err) => {
            this.err = err.response.data;
          });
        this.sending = false;
        this.clearForm();
      }, 2000);
    },
  },
  watch: {
    toggle() {
      app.lsPageToggle = this.toggle;
    },
    email(email) {
      this.loginData.email = email;
    },
    password(password) {
      this.loginData.password = password;
    },
  },
});

Vue.component("signup-comp", {
  template: `<div accept-charset="utf-8" style="width: 200px;">
                <nav style="text-align:center">
                    <ol class="breadcrumb-list">
                        <li class="breadcrumb-item" style="cursor: pointer" @click="toggle = !toggle">Вхід</li>
                        <li class="breadcrumb-item active">Реєстрація</li>
                    </ol>
                </nav>
                <br></br>
                <form novalidate @submit.prevent="validateUser">
                <div class="group">
                    <md-field :class="getValidationProfile('fname')">
                        <label>Ім’я</label>
                        <md-input v-model="profileData.fname" type="text" :disabled="sending"></md-input>
                        <span class="md-error" v-if="!$v.profileData.fname.required">Потрібно вказати ім’я</span>
                        <span class="md-error" v-else-if="!$v.profileData.fname.minLength">Недійсні дані</span>
                    </md-field>
                </div>
                <div class="group">
                    <md-field :class="getValidationProfile('lname')">
                        <label>Прізвище</label>
                        <md-input v-model="profileData.lname" type="text"></md-input>
                        <span class="md-error" v-if="!$v.profileData.lname.required">Потрібно вказати прізвище</span>
                        <span class="md-error" v-else-if="!$v.profileData.lname.minlength">Недійсні дані</span>
                    </md-field>
                </div>
                <div class="group">
                <md-field :class="getValidationProfile('gender')">
                    <label class="selected" for="gender">Стать</label>
                    <md-select v-model="profileData.gender" name="gender" id="gender" placeholder="Стать" md-dense :disabled="sending">
                        <md-option value="Male">Чоловік</md-option>
                        <md-option value="Female">Жінка</md-option>
                        <md-option value="Other">Інший варіант</md-option>
                    </md-select>
                    <span class="md-error">Вкажіть свою стать</span>
                </md-field>
                </div>
                <div class="group">
                      <md-datepicker id="date" v-model="profileData.birthday" :class="getValidationProfile('birthday')">
                        <label>Дата народження</label>
                        <span class="md-error">Вкажіть свою дату народження.</span>
                      </md-datepicker>
                </div>
                 <div class="group">
                    <md-field :class="getValidationSignUp('username')">
                        <label>Псевдонім</label>
                        <md-input v-model="signupData.username" type="text"></md-input>
                        <span class="md-error" v-if="!$v.signupData.username.required">Потрібно вказати псевдонім</span>
                    </md-field>
                </div>
                <div class="group">
                    <md-field :class="getValidationSignUp('email')">
                        <label>Електронна пошта</label>
                        <md-input v-model="signupData.email" type="text"></md-input>
                        <span class="md-error" v-if="!$v.signupData.email.required">Вкажіть електронну адресу</span>
                        <span class="md-error" v-if="!$v.signupData.email.email">Це не схоже на електронну адресу</span>
                    </md-field>
                </div>
                <div class="group">
                    <md-field :class="getValidationSignUp('password')">
                        <label>Пароль</label>
                        <md-input v-model="signupData.password" type="password"></md-input>
                        <span class="md-error" v-if="!$v.signupData.password.required">Потрібно вказати пароль</span>
                        <span class="md-error" v-else-if="!$v.signupData.password.minlength">Пароль повинен містити, як мінімум 8 символів.</span>
                    </md-field>
                </div>
                <md-progress-bar md-mode="indeterminate" v-if="sending" />
                <div v-if="success === false">
                <button type="submit" class="btn btn-outline-primary is-outlined" style="width:200px;" :disabled="sending">Зареєструватися</button>
                </div>
                <div v-else>
                    <button @click="toggle = !toggle" class="btn btn-outline-success is-outlined" style="width:200px;" :disabled="sending">Успіх</button>
                </div>
                </form>
                <div v-if="err != null" style="color:red; padding: 10px 0;">
                    <small>{{ err.msg }}</small>
                </div>
            </div>`,
  data: function () {
    return {
      signupData: {
        username: "",
        email: "",
        password: "",
      },
      profileData: {
        uuid: "",
        fname: "",
        lname: "",
        gender: "",
        // telno: "",
        birthday: null,
      },
      toggle: true,
      err: null,
      success: false,
      sending: false,
    };
  },
  validations: {
    profileData: {
      fname: {
        required,
        minLength: minLength(2),
      },
      lname: {
        required,
        minLength: minLength(3),
      },
      gender: {
        required,
      },
      birthday: {
        required,
      },
    },
    signupData: {
      email: {
        required,
        email,
      },
      password: {
        required,
        minLength: minLength(8),
      },
      username: { required },
    },
  },
  methods: {
    getValidationProfile(fieldName) {
      const field = this.$v.profileData[fieldName];

      if (field) {
        return {
          "md-invalid": field.$invalid && field.$dirty,
        };
      }
    },
    getValidationSignUp(fieldName) {
      const field = this.$v.signupData[fieldName];

      if (field) {
        return {
          "md-invalid": field.$invalid && field.$dirty,
        };
      }
    },
    clearForm: function () {
      this.$v.$reset();
      this.signupData.email = null;
      this.signupData.password = null;
      this.signupData.username = null;
      this.profileData.fname = null;
      this.profileData.lname = null;
      this.profileData.gender = null;
      this.profileData.birthday = null;
    },
    createUser: function () {
      this.sending = true;
      axios
        .post("/api/sign/up", this.signupData)
        .then((res) => {})
        .catch((err) => {
          this.err = err.response.data;
        });
      this.createProfile();
    },
    createProfile: function () {
      window.setTimeout(() => {
        axios
          .post("/api/client", this.profileData)
          .then((res) => {
            this.success = true;
            this.clearForm();
          })
          .catch((err) => {
            this.err = err.response.data;
          });
        this.sending = false;
      }, 5000);
    },
    validateUser() {
      this.$v.$touch();

      if (!this.$v.$invalid) {
        this.createUser();
      }
    },
  },
  watch: {
    toggle() {
      app.lsPageToggle = this.toggle;
    },
    username(username) {
      this.signupData.email = email;
    },
    password(password) {
      this.signupData.password = password;
    },
    fname(fname) {
      this.profileData.fname = fname;
    },
    lname(lname) {
      this.profileData.lname = lname;
    },
    gender(gender) {
      this.profileData.gender = gender;
    },
    birthday(birthday) {
      this.profileData.birthday = birthday;
    },
  },
});

// @click='view(info.id)'
Vue.component("product-comp", {
  template: ` <div v-show="info.id >= 0" :class="['rela-inline', 'product-card']" :key="info.id" :style="{'animation-delay':(info.delay*0.1)+'s'}">
                            <div class="rela-block product-pic" @click='view(info.id)' :style="{'background': 'url('+info.img+') center no-repeat'}">
                                                        <div class="product-hover text-white"><h2 class="abs-center">{{info.cost}}₴</h2></div>

                            </div>
                            <div class="rela-block product-info">
                                <div class="rela-block">
                                    <p>{{info.name}}</p>
                                    <p class="product-artist">{{info.artist}}</p>
                                </div>
                                <md-button class="vert-center product-cost md-icon-button" @click="addItem(info); toggle=true">
        <md-icon>add_shopping_cart</md-icon>
      </md-button>
                            </div>
                        </div>`,
  props: {
    info: {
      type: Object,
      default: {
        id: 0,
        name: "Untitled",
        artist: "Artist",
        desc: "Product description",
        cost: 0,
        medium: "",
        genre: "test",
        country: { String: "" },
        label: { String: "" },
        img: "https://picsum.photos/600/?random",
        qty: 0,
      },
    },
  },
  data: function () {
    return {
      toggle: false,
    };
  },
  methods: {
    view: function (id) {
      app.viewProduct(id);
    },
    addItem: function (product) {
      delete product["delay"];

      var index = 0;
      var exists = false;
      for (index; index < app.cart.length; ++index) {
        if (app.cart[index].name == product.name) {
          exists = true;
          break;
        }
      }
      if (exists) {
        app.incrementQty(index);
        app.menu;
      } else {
        product["qty"] = 1;
        app.cart.push(product);
        app.total += product.cost;
      }
    },
  },
  watch: {
    "info.qty": function () {
      if (info.qty === 0) {
        console.log("ZERO");
      }
    },
    toggle: function () {
      app.cartOpen = this.toggle;
      app.menuOpen = false;
    },
  },
});

Vue.component("comment-comp", {
  template: `
    <md-list class="md-double-line">
        <md-list-item>
            <md-avatar>
                <img :src="comment.img" alt="">
            </md-avatar>
            <div class="md-list-item-text">
                <span>{{comment.full_name}} </span>
                <span>пише "{{comment.text}}"</span>
            </div>
        </md-list-item>
    </md-list>`,
  props: {
    comment: {
      type: Object,
      default: {
        user_id: 0,
        product_id: 0,
        img: null,
        full_name: null,
        date: null,
        text: null,
      },
    },
  },
});

// - Vue Stuff -
var app = new Vue({
  el: "#app",
  data: {
    profile: [],
    comments: [],
    lsPageToggle: true,
    menuOpen: false,
    cartOpen: false,
    searchOpen: false,
    productViewOpen: false,
    filterOpen: false,
    profileOpen: false,
    checkoutOpen: false,
    loginStatus: false,
    editStatus: false,
    sending: false,
    currentViewedProduct: 0, // Product's id
    viewedProduct: {},
    searchInput: "",
    newItems: [],
    newItemPos: 0,
    products: products,
    filteredProducts: [],
    displayedProducts: [],
    displayPos: 0,
    genre: [],
    country: [],
    label: [],
    medium: [],
    year: [
      "< 1940",
      "1940-1950",
      "1950-1960",
      "1960-1970",
      "1970-1980",
      "1980-1990",
      "2000-2010",
      "2010-2021",
    ],
    currentGenre: "",
    currentCountry: "",
    currentYear: "",
    currentLabel: "",
    currentMedium: "",
    title: "",
    cart: [],
    order: null,
    cartIsEmpty: true,
    total: 0,
    err: "",
    comment: "",
  },
  watch: {
    cart: function () {
      if (this.cart.length) {
        this.cartIsEmpty = false;
      } else {
        this.cartIsEmpty = true;
      }
    },
    comment(comment) {
      this.comment = comment;
    },
    // cart["qty"]: function () {},
    searchInput: function () {
      this.searchText = this.searchInput;
      this.filteredProducts = [];
      this.currentGenre = "";
      this.currentLabel = "";
      this.currentYear = "";
      this.currentCountry = "";
      this.currentMedium = "";
      this.currentYear = "";

      // const options = {
      // includeScore: true,
      // shouldSort: true,
      // keys: ["artist", "name"],
      // };

      // var fuse = new Fuse(this.products, options);

      // if (this.searchText.length === 0) {
      //   this.updateFilteredProducts();
      // } else {
      //   var result = fuse.search(this.searchText);
      //   this.filteredProducts = result.map(function (a) {
      //     if (a.score < 0.5) {
      //       return a.item;
      //     }
      //   });
      // }
      // this.updateDisplayedProducts();

      if (this.searchText.length === 0) {
        this.updateFilteredProducts();
      } else {
        for (var product of this.products) {
          if (
            product.artist
              .toLowerCase()
              .includes(this.searchText.toLowerCase()) ||
            product.name.toLowerCase().includes(this.searchText.toLowerCase())
          ) {
            this.filteredProducts.push(product);
          }
        }
      }
      this.updateDisplayedProducts();
    },
    currentGenre: function () {
      if (this.currentGenre.length !== 0) {
        this.filterProductsByGenre();
      } else {
        this.updateFilteredProducts();
      }
    },
    currentCountry: function () {
      if (this.currentCountry.length !== 0) {
        this.filterProductsByCountry();
      } else {
        this.updateFilteredProducts();
      }
    },
    currentLabel: function () {
      if (this.currentLabel.length !== 0) {
        this.filterProductsByLabel();
      } else {
        this.updateFilteredProducts();
      }
    },
    currentMedium: function () {
      if (this.currentMedium.length !== 0) {
        this.filterProductsByMedium();
      } else {
        this.updateFilteredProducts();
      }
    },
    currentYear: function () {
      this.updateFilteredProducts();
    },
  },
  mounted: function () {
    axios.get("/api/genres").then((response) => {
      for (var i = 0; i < response.data.count; i++) {
        this.genre.push(response.data.genres[i]["name"]);
      }
    });

    this.medium = this.products.map((e) => {
      return e.medium;
    });
    this.medium = [...new Set(this.medium)];

    this.label = this.products.map((e) => {
      return e.label.String;
    });
    this.label = [...new Set(this.label)];
    this.label = this.label.filter(function (entry) {
      return entry.trim() != "";
    });

    this.country = this.products.map((e) => {
      return e.country.String;
    });
    this.country = [...new Set(this.country)];
    this.country = this.country.filter(function (entry) {
      return entry.trim() != "";
    });

    // axios.get("/api/profiles").then((response) => {
    //     this.profile.push(response.data[0]);
    // });
  },
  methods: {
    init: function () {
      this.products.unshift({
        id: -1,
        name: "Test",
        artist: "Test",
        desc: "",
        cost: 0,
        genre: "Test",
        country: { String: "" },
        label: { String: "" },
        medium: "",
        img: "https://picsum.photos/600/?random",
        date: "0",
      });
      app.updateNewItems();
      app.updateFilteredProducts();
    },
    updateNewItems: function () {
      // sort all of the products by date and then take the 10 newest
      this.newItems = [];
      var arr = [];
      // 1 because of the test product (Need to fix that)
      for (var i = 1; i < this.products.length; i++) {
        arr.push(this.products[i]);
      }
      arr.sort(function (a, b) {
        return new Date(b.date).getTime() - new Date(a.date).getTime();
      });

      for (var i = 0; i < 10 && i < arr.length; i++) {
        this.newItems.push(arr[i]);
      }
    },
    updateNewItemPos: function (num) {
      this.newItemPos += num;
      // Checks
      if (this.newItemPos < 0) {
        this.newItemPos = 0;
      }
      if (
        this.newItems.length > 1 &&
        this.newItemPos > this.newItems.length - 1
      ) {
        this.newItemPos = this.newItems.length - 1;
      }
    },
    updateFilteredProducts: function () {
      this.filteredProducts = [];

      // if (this.currentGenre === "Усі") {
      //   this.filteredProducts = this.products.filter((item) => item.genre);
      // } else {
      //   this.filteredProducts = this.products.filter((item) =>
      //     item.genre.includes(this.currentGenre)
      //   );
      // }

      this.filteredProducts = this.products;

      app.updateDisplayedProducts();
    },
    filterProductsByCountry: function () {
      this.filteredProducts = this.filteredProducts.filter((item) =>
        item.country.String.includes(this.currentCountry)
      );
      app.updateDisplayedProducts();
    },
    filterProductsByLabel: function () {
      this.filteredProducts = this.filteredProducts.filter((item) =>
        item.label.String.includes(this.currentLabel)
      );
      app.updateDisplayedProducts();
    },
    filterProductsByMedium: function () {
      this.filteredProducts = this.filteredProducts.filter((item) =>
        item.medium.includes(this.currentMedium)
      );
      app.updateDisplayedProducts();
    },
    filterProductsByGenre: function () {
      this.filteredProducts = this.filteredProducts.filter((item) =>
        item.genre.includes(this.currentGenre)
      );

      // for (var i in this.products) {
      //   if (this.products[i].genre == this.currentGenre) {
      //     this.filteredProducts.push(this.products[i]);
      //   }
      // }
      app.updateDisplayedProducts();
    },
    updateDisplayedProducts: function () {
      this.displayedProducts = [];
      console.log("updateDisplayedProducts");
      this.displayPos = 0;
      app.addDisplayedProducts();
    },
    addDisplayedProducts: function () {
      if (this.filteredProducts.length - this.displayPos <= 12) {
        this.displayedProducts = JSON.parse(
          JSON.stringify(this.filteredProducts)
        );
        for (var i = 0; i < this.displayedProducts.length; i++) {
          this.displayedProducts[i].delay = i - this.displayPos;
        }
        this.displayPos = this.filteredProducts.length;
      } else {
        // The ternary is for the test product. I really need to fix that.... :/
        for (var i = 0; i < (this.displayPos === 0 ? 13 : 12); i++) {
          this.displayedProducts.push(
            this.filteredProducts[i + this.displayPos]
          );
          this.displayedProducts[i + this.displayPos].delay = i;
        }
        this.displayPos = this.displayedProducts.length;
      }
    },
    updateViewedProduct: function () {
      this.viewedProduct = app.products.filter(function (el) {
        return el.id === app.currentViewedProduct;
      })[0];
      this.viewedProduct.country = this.viewedProduct.country.String;
      this.viewedProduct.label = this.viewedProduct.label.String;
    },
    viewProduct: function (id) {
      this.currentViewedProduct = id;
      app.updateViewedProduct();
      this.productViewOpen = true;
      axios.get(`/api/comment/${id}`).then((res) => {
        this.comments = res.data.comments;
      });
      // axios.get(`/api/user/${this.comments.user_id}`).then((res) => {
      // this.comments = res.data.review;
      // });
    },
    incrementQty: function (index) {
      this.total += this.cart[index].cost;
      this.cart[index].qty += 1;
    },
    decrementQty: function (index) {
      this.cart[index].qty -= 1;
      this.total -= this.cart[index].cost;
      if (this.cart[index].qty == 0) {
        this.cart.splice(index, 1);
      }
    },
    signOut: function () {
      axios
        .post(
          "/api/sign/out",
          {},
          {
            headers: {
              Authorization: `Bearer ${localStorage.access}`,
            },
          }
        )
        .then((res) => {})
        .catch((err) => {
          app.err = err;
          this.loginStatus = !this.loginStatus;
          localStorage.access = null;
          localStorage.refresh = null;
          localStorage.uuid = null;
          this.profile = [];
        });
    },
    createOrder: function () {
      this.sending = true;
      this.order = this.cart;
      axios
        .post(
          "/api/order",
          {
            user_id: localStorage.uuid,
            total: this.total,
          },
          {
            headers: {
              Authorization: `Bearer ${localStorage.access}`,
            },
          }
        )
        .then((res) => {})
        .catch((err) => {
          app.err = err;
        });
      window.setTimeout(() => {
        axios
          .get(`/api/order/${localStorage.uuid}`, {
            headers: {
              Authorization: `Bearer ${localStorage.access}`,
            },
          })
          .then((res) => {
            localStorage.order_id = res.data.order.order_id;
          })
          .catch((err) => {
            app.err = err;
          });
      }, 1000);
      window.setTimeout(() => {
        for (var i in this.order) {
          axios
            .post(
              "/api/order/list",
              {
                item_id: Number(i + 1),
                product_id: this.order[i].id,
                order_id: Number(localStorage.order_id),
                item_qty: this.order[i].qty,
                item_cost: this.order[i].cost,
                item_status: "Delivered",
              },
              {
                headers: {
                  Authorization: `Bearer ${localStorage.access}`,
                },
              }
            )
            .then((res) => {
              this.sending = false;
              this.checkoutOpen = false;
              this.cart = [];
              this.total = 0;
            })
            .catch((err) => {
              app.err = err;
            });
        }
      }, 5000);
    },
    addComment: function () {
      axios
        .post(
          "/api/review",
          {
            user_id: localStorage.uuid,
            product_id: this.viewedProduct.id,
            comment: this.comment,
          },
          {
            headers: {
              Authorization: `Bearer ${localStorage.access}`,
            },
          }
        )
        .then((res) => {
          console.log("Ok");
        })
        .catch((err) => {
          app.err = err;
        });
    },
    resetFilter: function () {
      this.currentYear = "";
      this.currentGenre = "";
      this.currentLabel = "";
      this.currentMedium = "";
      this.currentCountry = "";
      this.currentLabel = "";
    },
    // Counter: function (array) {
    //   array.forEach((item) => (item = (this[val] || 0) + 1));
    // },
  },
});
app.init();
// Scroll Function
// window.addEventListener('scroll', function() { app.pageScrolled = (window.scrollY > 0); }, false);

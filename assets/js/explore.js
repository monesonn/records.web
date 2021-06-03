// --- Notes ---
// https://dribbble.com/shots/4099945-Tacos-and-Trees-Marijuana-delivery-with-Tacos/attachments/939399
// 1039px is a good width

// - To Do -
// Finish designing product view
// Add Similar items to product view
// Finish the overall design as well
// *** Need to fix the animation with the test product and reuse of the first product and shit ***
// Make the search/filtering functions
// https://medium.com/@thierrymeier/filtering-and-sorting-best-practices-on-mobile-61626449cec
// Set up the cart system

// - Product Component -

Vue.component("login-comp", {
  template: `<div accept-charset="utf-8">
                <nav>
                    <ol class="breadcrumb-list">
                        <li class="breadcrumb-item active">Вхід</li>
                        <li class="breadcrumb-item" style="cursor: pointer" @click="app.lsPageToggle = !app.lsPageToggle">Реєстрація</li>
                    </ol>
                </nav>
                <br></br>
                <div class="group">
                    <input v-model="loginData.email" placeholder="" type="text" required>
                        <span class="highlight"></span>
                        <span class="bar"></span>
                        <label>Е-пошта</label>
                    </div>
                <div class="group">
                    <input v-model="loginData.password" type="password" required>
                        <span class="highlight"></span>
                        <span class="bar"></span>
                    <label>Пароль</label>
                </div>
                <button @click="checkLogin()" class="btn btn-outline-primary is-outlined" style="width:200px;">Увійти</button>
            </div>`,
  data: function () {
    return {
      loginData: {
        email: "",
        password: "",
      },
    };
  },
  // props: {
  // loginData: {
  // type: Object,
  // default: {
  // email: "client@mail.com",
  // password: "client",
  // },
  // },
  // },
  methods: {
    checkLogin: function () {
      axios
        .post("/api/sign/in", this.loginData)
        .then((res) => {
          localStorage.access = res.data.tokens.access;
          localStorage.refresh = res.data.tokens.refresh;
          app.loginStatus = true;
        })
        .catch((err) => {
          app.err = err;
        });
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

Vue.component("signup-comp", {
  template: `<div accept-charset="utf-8">
                <nav style="text-align:center">
                    <ol class="breadcrumb-list">
                        <li class="breadcrumb-item" style="cursor: pointer" @click="app.lsPageToggle = !app.lsPageToggle">Вхід</li>
                        <li class="breadcrumb-item active">Реєстрація</li>
                    </ol>
                </nav>
                <br></br>
                <div class="group">
                    <input v-model="signupData.firstName" type="text" required>
                        <span class="highlight"></span>
                        <span class="bar"></span>
                    <label>Ім’я</label>
                </div>
                <div class="group">
                    <input v-model="signupData.lastName" type="text" required>
                        <span class="highlight"></span>
                        <span class="bar"></span>
                    <label>Прізвище</label>
                </div>
                <div class="group">
                    <input v-model="signupData.gender" type="text" required>
                        <span class="highlight"></span>
                        <span class="bar"></span>
                    <label>Стать</label>
                </div>
                 <div class="group">
                    <input v-model="signupData.username" placeholder="" type="text" required>
                        <span class="highlight"></span>
                        <span class="bar"></span>
                        <label>Псевдонім</label>
                    </div>
                <div class="group">
                    <input v-model="signupData.password" type="password" required>
                        <span class="highlight"></span>
                        <span class="bar"></span>
                    <label>Пароль</label>
                </div>
                <button @click="" class="btn btn-outline-primary is-outlined" style="width:200px;">Зареєструватися</button>
            </div>`,
  data: function () {
    return {
      signupData: {
        username: "",
        password: "",
        firstName: "",
        lastName: "",
        gender: "",
        telNo: "",
        birthday: "",
      },
    };
  },
  methods: {
    signUp: function () {
      axios
        .post("/api/sign/up", this.signupData)
        .then((res) => {})
        .catch((err) => {
          app.err = err;
        });
    },
  },
});

// @click='view(info.id)'
Vue.component("product-comp", {
  template: ` <div v-show="info.id >= 0" :class="['rela-inline', 'product-card']" :key="info.id" :style="{'animation-delay':(info.delay*0.1)+'s'}">
                            <div class="rela-block product-pic" :style="{'background': 'url('+info.img+') center no-repeat'}">
                            <button class="btn btn-primary product-buy-button" @click="addItem(info);">До кошику <i class="bi bi-bag-plus"></i></button>
                            </div>
                            <div class="rela-block product-info">
                                <div class="rela-block">
                                    <p>{{info.name}}</p>
                                    <p class="product-artist">{{info.artist}}</p>
                                </div>
                                <div class="vert-center product-cost">{{info.cost}}₴</div>
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
        genre: "test",
        img: "https://picsum.photos/600/?random",
        qty: 0,
      },
    },
  },
  methods: {
    view: function (id) {
      app.viewProduct(id);
    },
    addItem: function (product) {
      app.total += product.cost;
      localStorage.total = app.total;
      product["qty"] = 1;
      delete product["delay"];
      app.cart.push(product);
      console.log(app.cart);
    },
  },
  watch: {
    "info.qty": function () {
      if (info.qty === 0) {
        console.log("ZERO");
      }
    },
  },
});

// - Vue Stuff -
var app = new Vue({
  el: "#app",
  data: {
    profile: [],
    lsPageToggle: true,
    menuOpen: false,
    cartOpen: false,
    searchOpen: false,
    productViewOpen: false,
    profileOpen: false,
    checkoutOpen: false,
    loginStatus: false,
    currentViewedProduct: 0, // Product's id
    viewedProduct: {},
    searchInput: "",
    newItems: [],
    newItemPos: 0,
    products: products,
    filteredProducts: [],
    displayedProducts: [],
    displayPos: 0,
    genre: ["Усі"],
    currentGenre: "Усі",
    title: "",
    cart: [],
    cartIsEmpty: true,
    total: 0,
    err: "",
  },
  watch: {
    cart: function () {
      if (this.cart.length) {
        this.cartIsEmpty = false;
      } else {
        this.cartIsEmpty = true;
      }
    },
    searchInput: function () {
      this.searchText = this.searchInput;
      this.filteredProducts = [];

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
  },
  mounted: function () {
    axios.get("/api/genres").then((response) => {
      for (var i = 0; i < response.data.count; i++) {
        this.genre.push(response.data.genres[i]["name"]);
      }
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
      for (var i in this.products) {
        if (
          this.products[i].genre === this.currentGenre ||
          this.currentGenre === "Усі"
        ) {
          this.filteredProducts.push(this.products[i]);
        }
      }
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
    },
    viewProduct: function (id) {
      this.currentViewedProduct = id;
      app.updateViewedProduct();
      this.productViewOpen = true;
    },
    incrementQty: function (index) {
      this.total += this.cart[index].cost;
      this.cart[index].qty += 1;
    },
    decrementQty: function (index) {
      this.cart[index].qty -= 1;
      if (this.cart[index].qty == 0) {
        this.cart.pop();
      } else {
        this.total -= this.cart[index].cost;
      }
    },
    // Counter: function (array) {
    //   array.forEach((item) => (item = (this[val] || 0) + 1));
    // },
  },
});
app.init();
// Scroll Function
// window.addEventListener('scroll', function() { app.pageScrolled = (window.scrollY > 0); }, false);

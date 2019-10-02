(function() {
  'use strict';
  
  angular
  .module('blog-app')
  .controller('app-controller', ctrl1);

  function ctrl1($route, $http, $window){
    
    /*jshint validthis: true*/
    var vm = this;
    vm.post = {};
    vm.postToEdit = {};
    vm.posts = [];
    vm.getPosts = getPosts;
    vm.editPost = editPost;
    vm.deletePost = deletePost;
    vm.createPost = createPost;
    vm.setPost = setPost;
    
    function setPost(post){
      vm.postToEdit = post;
    }
    
    function redirectHome(){
      $window.location.href = '#!/posts';
    }

    function getPosts() {
      var promise = $http.get('http://34.73.234.35:3000/posts');
      promise.then(successCallback, failureCallback)
      function successCallback(result) {
        vm.posts = result.data;
        setPost({});
        vm.post = {};
      }
      function failureCallback(result) {
        console.log("Error", result)
      }
    }

    function editPost() {
      var req = {
         method: 'PUT',
         url: 'http://34.73.234.35:3000/posts',
         headers: {
           'Content-Type': "application/json"
         },
         data: {
          id: parseInt($route.current.params.id),
          title: vm.postToEdit.title,
          body: vm.postToEdit.body
        }
      }
      $http(req).then(function(){
        console.log("Data updated");
        setPost({});
        redirectHome();
      }, function(){
        console.log("There was an error");
      });
    }
    
    function deletePost() {
      var promise = $http.delete('http://34.73.234.35:3000/posts?id=' + parseInt($route.current.params.id));
      promise.then(successCallback, failureCallback)
      function successCallback(result) {
        console.log("Data deleted");
        setPost({});
        redirectHome();
      }
      function failureCallback(result) {
        console.log("Error", result)
      }
    }

    function createPost() {
      var req = {
         method: 'POST',
         url: 'http://34.73.234.35:3000/posts',
         headers: {
           'Content-Type': "application/json"
         },
         data: {
          title: vm.post.title,
          body: vm.post.body
        }
      }
      $http(req).then(function(){
        console.log("Data created");
        vm.post = {};
        redirectHome();
      }, function(){
        console.log("There was an error");
      });
    }
    
 } //End Controller Function
  
})();
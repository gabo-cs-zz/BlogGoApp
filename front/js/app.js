(function() {
  'use strict';
  
  angular
    .module('blog-app', ['ngRoute'])
    .config(['$routeProvider', function($routeProvider){
      $routeProvider
        .when('/', {
        templateUrl: 'views/dashboard.html'
      }).when('/posts', {
        templateUrl: 'views/posts.html'
      }).when('/create', {
        templateUrl: 'views/create_post.html'
      }).when('/posts/:id', {
        templateUrl: 'views/edit_post.html'
      }).otherwise({
        redirectTo: '/posts'
      });
    }])
    .filter('trusted', ['$sce', function ($sce) {
      return $sce.trustAsResourceUrl;
    }])
    .directive('ngConfirmClick', [
      function(){
        return {
          link: function (scope, element, attr) {
            var msg = attr.ngConfirmClick || "Are you sure?";
            var clickAction = attr.confirmedClick;
            element.bind('click',function (event) {
              if ( window.confirm(msg) ) {
                scope.$eval(clickAction)
              }
            });
          }
        };
    }]);
  
})();
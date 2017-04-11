angular
    .module('TodoApp', [])
    .directive('focusMe', function ($timeout) {
        return {
            scope: {trigger: '@focusMe'},
            link: function (scope, element) {
                scope.$watch('trigger', function (value) {
                    if (value === "true") {
                        $timeout(function () {
                            element[0].focus();
                        });
                    }
                });
            }
        };
    })
    .controller('AppCtrl', ['$http', function AppCtrl($http) {

        // ------------- Public -------------
        vm = this;
        vm.tasks = [];
        vm.working = false;
        vm.focused = true;

        vm.addTask = addTask;
        vm.removeTask = removeTask;
        vm.toggleDone = toggleDone;


        // ------------- Private -------------

        var logError = function (data, status) {
            alert('code ' + status + ': ' + data);
            console.log('code ' + status + ': ' + data);
            vm.working = false;
        };

        var refresh = function () {
            return $http.get('/task/').success(function (data) {
                vm.tasks = data.Tasks;
                vm.working = false;
                vm.todoText = '';
            }).error(logError);
        };

        function addTask() {
            vm.working = true;
            $http.post('/task/', {Title: vm.todoText})
                .error(logError)
                .success(function () {
                refresh();
            });
        }

        function removeTask(task) {
            vm.working = true;
            $http.delete('/task/' + task.ID).error(logError).success(function () {
                refresh();
            });
        }

        function toggleDone(task) {
            var data = {ID: task.ID, Title: task.Title, Done: !task.Done}
            $http.put('/task/' + task.ID, data).error(logError).success(function () {
                task.Done = !task.Done
            });
        }

        refresh();
    }]);

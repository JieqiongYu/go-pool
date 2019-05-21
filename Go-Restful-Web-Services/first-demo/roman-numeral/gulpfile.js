// 1. Import Gulp and Gulp-shell
var gulp = require("gulp");
var shell = require('gulp-shell');

// 2. Create tasks with `shell.task` as the function to execute. 
// 3. `shell.task` can execute a command-line instruction. Keep your shell commands inside that function. 
// This compiles new binary with source change
gulp.task("install-binary", shell.task([
  'go install github.com/restful-web-services/roman-numeral/romanserver'
]));


// Second argument tells install-binary is a deapendency for restart-supervisor
gulp.task("restart-supervisor",  shell.task([
  'supervisorctl restart romanserver'
]));
gulp.task(gulp.series(gulp.parallel("restart-supervisor", "install-binary")),  shell.task([
  'supervisorctl restart romanserver'
]))

gulp.task('watch', function() {
  // 4. Add a watch task for watching source files. The task list will be executed when files are modified. 
  // Watch the source code for all changes
  gulp.watch("*", gulp.series(gulp.parallel('install-binary', 'restart-supervisor')));

});

// 5. Create a default task for running. Add a watch to it. 
gulp.task('default', gulp.series(gulp.parallel('watch')));

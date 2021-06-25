// this is the function that is used in the experiment

function mean(count) {
  total = 0 ;
  for (i = 0; i < count; i++) {
    total = total + i ;
  }
  mean = total/count ;
  return mean ;
}

mean(50);

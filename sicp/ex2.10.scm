
(define (make-interval a b)
  if (< a b) (cons a b) (cons b a))

(define (lower-bound x)
  (car x))

(define (upper-bound x)
  (cdr x))

(define (add-interval a b)
  (make-interval
    (+ (lower-bound a) (lower-bound b))
    (+ (upper-bound a) (upper-bound b))))

(define (mul-interval a b)
  (let (
        (p1 (* (lower-bound a) (lower-bound b)))
        (p2 (* (lower-bound a) (upper-bound b)))
        (p3 (* (upper-bound a) (lower-bound a)))
        (p4 (* (upper-bound a) (upper-bound b))))
      (make-interval (min p1 p2 p3 p4) (max p1 p2 p3 p4))))

(define (div-interval a b)
  (if (or (= (lower-bound b) 0) (= (upper-bound b) 0))
    (error "error")
    (mul-interval
      a
      (make-interval (/ 1.0 (upper-bound b))
                     (/ 1.0 (lower-bound b))))))

(define (sub-interval a b)
  (make-interval
    (- (lower-bound a) (upper-bound b))
    (- (upper-bound a) (lower-bound b))))

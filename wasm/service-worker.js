var cacheName = 'opencc'

var filesToCache = [
  './',
  './index.html',
  './wasm_exec.js',
  './opencc.wasm',
  'https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.0/css/bootstrap.min.css',
  'https://cdnjs.cloudflare.com/ajax/libs/vue/2.6.11/vue.min.js'
]

self.addEventListener('install', function (e) {
  e.waitUntil(
    caches.open(cacheName).then(function (cache) {
      return cache.addAll(filesToCache)
    })
  )
})

self.addEventListener('fetch', function (e) {
  e.respondWith(
    caches.match(e.request).then(function (response) {
      return fetch(e.request).then(res => {
        return caches.open(cacheName).then(function (cache) {
          return cache.put(e.request, res.clone()).then(function () {
            return res
          })
        })
      }).catch(() => {
        return response
      })
    })
  )
})

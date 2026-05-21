Rails.application.routes.draw do
  post '/register', to: 'auth#register'
  post '/login', to: 'auth#login'
  get '/me', to: 'auth#me'
  post '/validate', to: 'auth#validate'
end

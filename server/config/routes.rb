Rails.application.routes.draw do
  get 'user_plans/index'

  namespace :api, format: 'json' do
    namespace :v1 do
      post 'login' => 'user_token#create'
      post 'signup' => 'user#create'
      get 'ping' => 'authenticated#ping'
      # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
      resource :user, only: [:show, :update, :destroy] do
        resources :plans
      end
      namespace :admin do
        resources :users, only: [:index] do
          resources :plans, only: [:show, :update, :destroy]
          resources :user_plans, only: [:index]
        end
        resources :plans, only: [:index]
      end
    end
  end
end

Rails.application.routes.draw do
  namespace :api, format: 'json' do
    namespace :v1 do
      post 'login' => 'user_token#create'
      post 'signup' => 'user#create'
      # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
      resource :user, only: [:show, :update, :destroy] do
        resources :plans
      end
      namespace :admin do
        resources :users, only: [:index] do
          resources :plans, only: [:show, :update, :destroy]
        end
        resources :plans, only: [:index]
      end
    end
  end
end
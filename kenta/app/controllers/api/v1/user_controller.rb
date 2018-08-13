class Api::V1::UserController < ApplicationController
  before_action :authenticate_user, only: %i[show update destroy]
  before_action :load_resource

  def show
    render json: @user
  end

  #signup処理ここでやってる
  def create
    user = User.new(user_params)
    if user.save
      auth_token = Knock::AuthToken.new payload: { sub: user.id }
      render json: auth_token, status: :created
    else
      render json: { error: user.errors.full_messages }, status: :unprocessable_entity
    end
  end

  def update
    if @user.update!(user_params)
      render json: @user, status: 200
    else
      render json: { errors: @user.errors }, status:422
    end
  end

  def destroy
    @user.destroy
    head 204
  end

  private

  def load_resource
    case params[:action].to_sym
      when :show, :update
        @user = User.select(:name, :sex, :email, :birthday, :university).find(current_user.id)
      when :destroy
        @user = User.find(current_user.id)
      end
  end

  def user_params
    params.permit(:name, :sex, :email, :password, :password_confirmation, :birthday, :university)
  end
end

class Api::V1::AuthenticatedController < ApplicationController
  def ping
    if current_user
      render json: { message: "ログイン中です" }, status: 200
    else
      render json: { message: "ログアウトしてます" }, status: 401
    end
  end
end

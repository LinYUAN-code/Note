use snake_game::snake::SnakeGame;

mod snake_game;
fn main() {
    let mut snake_game = SnakeGame::new(20, 20);
    snake_game.game_start();
}

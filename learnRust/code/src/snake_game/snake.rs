use piston_window::{
    clear, types::Color, Button, Context, G2d, Key, PistonWindow, PressEvent, UpdateEvent,
    WindowSettings,
};
use rand::Rng;

use crate::snake_game::draw::{BACK_COLOR, SNAKE_BODY_COLOR};

use super::draw::{draw_block, to_coord, FAIL_BACK_COLOR, FOOD_COLOR};

pub type Pos = [i32; 2];
// 多少秒移动一次
const MOVING_PERIOD: f64 = 0.1;
struct Food {
    pos: Pos,
    width: u32,
    height: u32,
}
impl Food {
    pub fn new(width: u32, height: u32) -> Food {
        let x = rand::thread_rng().gen_range(0..width) as i32;
        let y = rand::thread_rng().gen_range(0..height) as i32;
        Food {
            pos: [x, y],
            width,
            height,
        }
    }
    pub fn random_pos(&mut self) {
        let x = rand::thread_rng().gen_range(0..self.width) as i32;
        let y = rand::thread_rng().gen_range(0..self.height) as i32;
        self.pos = [x, y]
    }
    pub fn update_to_board(&self, board: &mut Vec<Vec<Element>>) {
        board[self.pos[0] as usize][self.pos[1] as usize] = Element::Food(FOOD_COLOR);
    }
    pub fn handle_input(&mut self, key: Key) {}
    pub fn update(&mut self) {}
}
enum SnakeNodeKind {
    Normal,
}
struct SnakeNode {
    pos: Pos,
    kind: SnakeNodeKind,
    dir: Direction,
}
impl SnakeNode {
    pub fn new(x: i32, y: i32, kind: SnakeNodeKind, dir: Direction) -> SnakeNode {
        SnakeNode {
            pos: [x, y],
            kind,
            dir,
        }
    }
    pub fn update_to_board(&self, board: &mut Vec<Vec<Element>>) {
        board[self.pos[0] as usize][self.pos[1] as usize] = Element::SnakeBody(SNAKE_BODY_COLOR);
    }
    fn update(&mut self) {
        match self.dir {
            Direction::Up => self.pos[1] -= 1,
            Direction::Down => self.pos[1] += 1,
            Direction::Left => self.pos[0] -= 1,
            Direction::Right => self.pos[0] += 1,
        }
    }
    // 生一个节点出来
    pub fn give_birth(&self) -> SnakeNode {
        let mut snake_node = SnakeNode {
            pos: self.pos.clone(),
            kind: SnakeNodeKind::Normal,
            dir: self.dir.clone(),
        };
        match self.dir {
            Direction::Up => snake_node.pos[1] += 1,
            Direction::Down => snake_node.pos[1] -= 1,
            Direction::Left => snake_node.pos[0] += 1,
            Direction::Right => snake_node.pos[0] -= 1,
        }
        snake_node
    }
}
#[derive(PartialEq)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

impl Copy for Direction {}
impl Clone for Direction {
    fn clone(&self) -> Self {
        match self {
            Self::Up => Self::Up,
            Self::Down => Self::Down,
            Self::Left => Self::Left,
            Self::Right => Self::Right,
        }
    }
}
struct Snake {
    snake_body: Vec<SnakeNode>,
}
impl Snake {
    pub fn new() -> Snake {
        let mut snake_body = Vec::new();
        snake_body.push(SnakeNode::new(
            3,
            2,
            SnakeNodeKind::Normal,
            Direction::Right,
        ));
        snake_body.push(SnakeNode::new(
            2,
            2,
            SnakeNodeKind::Normal,
            Direction::Right,
        ));
        snake_body.push(SnakeNode::new(
            1,
            2,
            SnakeNodeKind::Normal,
            Direction::Right,
        ));
        Snake { snake_body }
    }
    pub fn update_to_board(&self, board: &mut Vec<Vec<Element>>) {
        for snake_node in self.snake_body.iter() {
            snake_node.update_to_board(board);
        }
    }

    pub fn handle_input(&mut self, key: Key) {
        match key {
            Key::Up => self.change_dir(Direction::Up, true),
            Key::Down => self.change_dir(Direction::Down, true),
            Key::Left => self.change_dir(Direction::Left, true),
            Key::Right => self.change_dir(Direction::Right, true),
            _ => (),
        };
    }
    fn change_dir(&mut self, dir: Direction, isOnlyHead: bool) {
        let head_dir = self.snake_body[0].dir;

        if dir == Direction::Up && head_dir == Direction::Down {
            return;
        }
        if dir == Direction::Down && head_dir == Direction::Up {
            return;
        }
        if dir == Direction::Left && head_dir == Direction::Right {
            return;
        }
        if dir == Direction::Right && head_dir == Direction::Left {
            return;
        }
        if !isOnlyHead {
            for i in (1..self.snake_body.len()).rev() {
                self.snake_body[i].dir = self.snake_body[i - 1].dir;
            }
        }
        self.snake_body[0].dir = dir
    }
    pub fn update(&mut self, food: &mut Food) {
        for snake_node in self.snake_body.iter_mut() {
            snake_node.update();
        }
        if self.snake_body[0].pos == food.pos {
            food.random_pos();
            let new_snake_node = self.snake_body.last().unwrap().give_birth();
            self.snake_body.push(new_snake_node);
        }
        self.change_dir(self.snake_body[0].dir, false);
    }
}
pub struct SnakeGame {
    window: Option<PistonWindow>,
    width: u32,
    height: u32,
    game_board: Vec<Vec<Element>>,
    snake: Snake,
    food: Food,
    game_state: GameState,
    waitting_time: f64,
}
#[derive(PartialEq)]
pub enum GameState {
    Over,
    Continue,
}
pub enum Element {
    Empty(Color),
    SnakeBody(Color),
    Food(Color),
    Wall(Color),
}

impl SnakeGame {
    pub fn new(width: u32, height: u32) -> SnakeGame {
        // 初始化棋盘
        let mut game_board = Vec::new();
        for _ in 0..width {
            let mut row = Vec::new();
            for _ in 0..height {
                row.push(Element::Empty(BACK_COLOR))
            }
            game_board.push(row)
        }
        // 初始化蛇
        let snake = Snake::new();
        // 初始化食物
        let food = Food::new(width, height);
        SnakeGame {
            window: None,
            width,
            height,
            game_board,
            snake,
            food,
            game_state: GameState::Continue,
            waitting_time: 0f64,
        }
    }
    pub fn game_start(&mut self) {
        println!("game_start");
        // 初始化参数
        let mut window_settings = WindowSettings::new(
            "Lin Yuan Snake Game",
            [to_coord(self.width), to_coord(self.height)],
        )
        .exit_on_esc(true);

        // Fix vsync extension error for linux
        window_settings.set_vsync(true);
        let window: PistonWindow = window_settings.build().unwrap();

        self.window = Some(window);

        self.run();
    }
    pub fn run(&mut self) {
        let window = self.window.as_mut().unwrap();
        while let Some(event) = window.next() {
            // 生成game_board
            if self.game_state == GameState::Over {
                clear_game_board(&mut self.game_board, FAIL_BACK_COLOR);
                window.draw_2d(&event, |c, g, _| {
                    // 绘制
                    draw(&self.game_board, c, g);
                });
                continue;
            } else {
                clear_game_board(&mut self.game_board, BACK_COLOR);
            }
            self.food.update_to_board(&mut self.game_board);
            self.snake.update_to_board(&mut self.game_board);
            window.draw_2d(&event, |c, g, _| {
                // 绘制
                draw(&self.game_board, c, g);
            });
            // 游戏更新--蛇移动以及其他逻辑(如：食物超时跳转位置)
            event.update(|args| {
                self.waitting_time += args.dt;
                if self.waitting_time >= MOVING_PERIOD {
                    self.snake.update(&mut self.food);
                    self.food.update();
                    self.game_state =
                        get_game_state(&self.snake, &self.game_board, self.width, self.height);
                    self.waitting_time = 0f64;
                }
            });

            // 处理输入
            if let Some(Button::Keyboard(key)) = event.press_args() {
                self.food.handle_input(key);
                self.snake.handle_input(key);
            }
        }
    }
}
fn get_game_state(
    snake: &Snake,
    game_board: &Vec<Vec<Element>>,
    width: u32,
    height: u32,
) -> GameState {
    let x = snake.snake_body[0].pos[0];
    let y = snake.snake_body[0].pos[1];
    if x < 0 || x >= width as i32 || y < 0 || y >= height as i32 {
        return GameState::Over;
    }
    match game_board[x as usize][y as usize] {
        Element::SnakeBody(_) => GameState::Over,
        Element::Wall(_) => GameState::Over,
        _ => GameState::Continue,
    }
}
pub fn clear_game_board(game_board: &mut Vec<Vec<Element>>, color: Color) {
    let n = game_board.len();
    let m = game_board[0].len();
    for i in 0..n {
        for j in 0..m {
            game_board[i][j] = Element::Empty(color);
        }
    }
}
pub fn draw(game_board: &Vec<Vec<Element>>, c: Context, g: &mut G2d) {
    clear(BACK_COLOR, g);
    let n = game_board.len();
    let m = game_board[0].len();
    for i in 0..n {
        for j in 0..m {
            match game_board[i][j] {
                Element::SnakeBody(color) => draw_block(color, i as u32, j as u32, &c, g),
                Element::Empty(color) => draw_block(color, i as u32, j as u32, &c, g),
                Element::Food(color) => draw_block(color, i as u32, j as u32, &c, g),
                Element::Wall(color) => draw_block(color, i as u32, j as u32, &c, g),
            }
        }
    }
}

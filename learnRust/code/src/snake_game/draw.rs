use piston_window::{rectangle, types::Color, Context, G2d, PistonWindow, WindowSettings};

const BLOCK_SIZE: f64 = 25.0;
pub const BACK_COLOR: Color = [0.204, 0.286, 0.369, 1.0];
pub const FAIL_BACK_COLOR: Color = [0.904, 0.286, 0.369, 0.4];
pub const SNAKE_BODY_COLOR: Color = [0.18, 0.80, 0.44, 1.0];
pub const FOOD_COLOR: Color = [0.88, 0.50, 0.54, 1.0];

pub fn draw_block(color: Color, x: u32, y: u32, con: &Context, g: &mut G2d) {
    let gui_x = to_coord(x);
    let gui_y = to_coord(y);

    rectangle(
        color,
        [gui_x, gui_y, BLOCK_SIZE, BLOCK_SIZE],
        con.transform,
        g,
    );
}

pub fn to_coord(len: u32) -> f64 {
    (len as f64) * BLOCK_SIZE
}

use crypto::digest::Digest;
use crypto::md5::Md5;
use hex::ToHex;

fn main() {
    let input = "ckczppom".to_owned();

    let mut five_zeroes = 0;
    let mut i = 0;

    let mut hash = Md5::new();

    while five_zeroes == 0 {
        i += 1;

        hash.input(input.as_bytes());
        hash.input(i.to_string().as_bytes());

        let mut output = [0; 16]; // An MD5 is 16 bytes
        hash.result(&mut output);

        if output[0] == 0 && output[1] == 0 && output[2] < 16
        {
            five_zeroes = i;
        }

        hash.reset();
    }

    println!("{:?}", five_zeroes);
}
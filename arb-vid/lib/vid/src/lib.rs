extern crate ark_bls12_381;
extern crate jf_primitives;

use ark_bls12_381::Bls12_381;
use jf_primitives::{
    pcs::{checked_fft_size, prelude::UnivariateKzgPCS, PolynomialCommitmentScheme},
    vid::advz::Advz,
};
use std::ffi::CStr;

#[no_mangle]
pub extern "C" fn mock_crypto(msg: *const libc::c_char) {
    let msg_cstr = unsafe { CStr::from_ptr(msg) };
    let msg = msg_cstr.to_str().unwrap();
    let _advz: Advz<Bls12_381, sha2::Sha256>;
    let (payload_chunk_size, num_storage_nodes) = (8, 10);

    let mut rng = jf_utils::test_rng();
    let srs = UnivariateKzgPCS::<Bls12_381>::gen_srs_for_testing(
        &mut rng,
        checked_fft_size(payload_chunk_size - 1).unwrap(),
    )
    .unwrap();
    _advz = Advz::new(payload_chunk_size, num_storage_nodes, srs).unwrap();
    println!("({})", msg)
}

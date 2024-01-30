use ark_bls12_381::Bls12_381;
use go_abi::*;
use jf_primitives::{
    pcs::{checked_fft_size, prelude::UnivariateKzgPCS, PolynomialCommitmentScheme},
    vid::advz::{payload_prover::LargeRangeProof, Advz},
};

// extern "C" {
//     pub fn BrotliDecoderDecompress(
//         encoded_size: usize,
//         encoded_buffer: *const u8,
//         decoded_size: *mut usize,
//         decoded_buffer: *mut u8,
//     ) -> u32;

//     pub fn BrotliEncoderCompress(
//         quality: u32,
//         lgwin: u32,
//         mode: u32,
//         input_size: usize,
//         input_buffer: *const u8,
//         encoded_size: *mut usize,
//         encoded_buffer: *mut u8,
//     ) -> u32;
// }

// const BROTLI_MODE_GENERIC: u32 = 0;
// const BROTLI_RES_SUCCESS: u32 = 1;

#[no_mangle]
pub unsafe extern "C" fn go__github_com_offchainlabs_nitro_arbcompress_verifyNamespace(
    sp: GoStack,
) {
    // let advz: Advz<Bls12_381, sha2::Sha256>;
    // let (payload_chunk_size, num_storage_nodes) = (8, 10);

    // let mut rng = jf_utils::test_rng();
    // let srs = UnivariateKzgPCS::<Bls12_381>::gen_srs_for_testing(
    //     &mut rng,
    //     checked_fft_size(payload_chunk_size - 1).unwrap(),
    // )
    // .unwrap();
    // advz = Advz::new(payload_chunk_size, num_storage_nodes, srs).unwrap();
    return unimplemented!("asdf");
}

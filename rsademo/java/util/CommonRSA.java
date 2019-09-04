package util;


import org.apache.commons.codec.binary.Base64;

import javax.crypto.Cipher;
import java.io.UnsupportedEncodingException;
import java.security.KeyFactory;
import java.security.NoSuchAlgorithmException;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.Signature;
import java.security.spec.InvalidKeySpecException;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;
import java.util.logging.Logger;

/**
 * Created by 747endetta on 3/11/16.
 */
public class CommonRSA {
	protected static final String CHARACTER_SET = "utf-8";
	protected static final String ALGORITHM = "RSA";
	protected static final String SIGN_ALGORITHMS = "SHA1WithRSA";
	protected static final String RSA256_SIGN_ALGORITHM = "SHA256WithRSA";
	protected static final String DEFAULT_CIPHER = "RSA/ECB/PKCS1Padding";


	public static String sign(String content, String privateKey) {
		try {
			KeyFactory keyFactory = KeyFactory.getInstance(ALGORITHM);
			PKCS8EncodedKeySpec pkcs8EncodedKeySpec = new PKCS8EncodedKeySpec(decodeBASE64(privateKey));
			PrivateKey priKey = keyFactory.generatePrivate(pkcs8EncodedKeySpec);

			Signature signature = Signature.getInstance(SIGN_ALGORITHMS);
			signature.initSign(priKey);
			signature.update(content.getBytes(CHARACTER_SET));
			byte[] signed = signature.sign();
			return encodeBASE64(signed);
		} catch (Exception e) {
			e.printStackTrace();
		}
		return null;
	}

	public static String rsa256Sign(String content, String privateKey) {
		try {
			KeyFactory keyFactory = KeyFactory.getInstance(ALGORITHM);
			PKCS8EncodedKeySpec pkcs8EncodedKeySpec = new PKCS8EncodedKeySpec(decodeBASE64(privateKey));
			PrivateKey priKey = keyFactory.generatePrivate(pkcs8EncodedKeySpec);

			Signature signature = Signature.getInstance(RSA256_SIGN_ALGORITHM);
			signature.initSign(priKey);
			signature.update(content.getBytes(CHARACTER_SET));
			byte[] signed = signature.sign();
			return encodeBASE64(signed);
		} catch (Exception e) {
		}
		return null;
	}

	public static boolean doCheck(String content, String sign, String publicKey) {
		try {
			PublicKey pubKey = parsePublicKey(publicKey);

			Signature signature = Signature.getInstance(SIGN_ALGORITHMS);
			signature.initVerify(pubKey);
			signature.update(content.getBytes(CHARACTER_SET));
			return signature.verify(decodeBASE64(sign));
		} catch (Exception e) {
			e.printStackTrace();
		}

		return false;
	}

	public static boolean doCheckRSA256(String content, String sign, String publicKey) {
		try {
			PublicKey pubKey = parsePublicKey(publicKey);

			Signature signature = Signature.getInstance(RSA256_SIGN_ALGORITHM);
			signature.initVerify(pubKey);
			signature.update(content.getBytes(CHARACTER_SET));
			return signature.verify(decodeBASE64(sign));
		} catch (Exception e) {
		}

		return false;
	}

	public static PublicKey parsePublicKey(String publicKey) {
		try {
			KeyFactory keyFactory = KeyFactory.getInstance(ALGORITHM);
			byte[] encodedKey = decodeBASE64(publicKey);
			return keyFactory.generatePublic(new X509EncodedKeySpec(encodedKey));
		} catch (NoSuchAlgorithmException | InvalidKeySpecException | UnsupportedEncodingException e) {
		}
		return null;
	}

	public static String encrypt(String content, String publicKey) {
		return encrypt(content, publicKey, DEFAULT_CIPHER);
	}

	public static String encrypt(String content, String publicKey, String c) {
		try {
			KeyFactory keyFactory = KeyFactory.getInstance(ALGORITHM);
			byte[] encodedKey = decodeBASE64(publicKey);
			PublicKey pubKey = keyFactory.generatePublic(new X509EncodedKeySpec(encodedKey));
			Cipher cipher = Cipher.getInstance(c);
			cipher.init(Cipher.ENCRYPT_MODE, pubKey);
			return encodeBASE64(cipher.doFinal(content.getBytes()));
		} catch (Exception e) {
			e.printStackTrace();
		}
		return null;


	}

	public static String decrypt(String content, String privateKey) {
		return decrypt(content, privateKey, DEFAULT_CIPHER);
	}
	public static String decrypt(String content, String privateKey, String c) {
		try {
			KeyFactory keyFactory = KeyFactory.getInstance(ALGORITHM);
			PKCS8EncodedKeySpec pkcs8EncodedKeySpec = new PKCS8EncodedKeySpec(decodeBASE64(privateKey));
			PrivateKey priKey = keyFactory.generatePrivate(pkcs8EncodedKeySpec);
			Cipher cipher = Cipher.getInstance(c);
			cipher.init(Cipher.DECRYPT_MODE, priKey);
			return new String(cipher.doFinal(decodeBASE64(content)));
		} catch (Exception e) {
			e.printStackTrace();
		}
		return null;

	}

	private static String encodeBASE64(byte[] value) {
		return new String(Base64.encodeBase64(value));
	}

	private static byte[] decodeBASE64(String value) throws UnsupportedEncodingException {
		return Base64.decodeBase64(value.getBytes(CHARACTER_SET));
	}
}

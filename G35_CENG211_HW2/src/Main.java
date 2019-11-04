import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.StringTokenizer;

import com.sun.org.apache.bcel.internal.generic.LNEG;

public class Main {

	public static void main(String[] args) throws IOException {
		File csv = new File("CENG211_HW2_ArtVaultData.csv");
		BufferedReader br1 = new BufferedReader(new FileReader(csv));
		String line = br1.readLine();
		ArrayList<String> arraylist1 = new ArrayList<String>();
		while(line != null) {
			String delims = ",";
			StringTokenizer st1 = new StringTokenizer(line,delims);
			while(st1.hasMoreTokens()) {
				arraylist1.add(st1.nextToken());
				}
			line = br1.readLine();
			
			}
		
				
	}
		
		
			
}
	




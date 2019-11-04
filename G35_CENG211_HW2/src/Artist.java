
public class Artist extends Person{
	private String placeOfResidence;

	public Artist(String name,String birtdate,String deathdate,String country,String placeOfResidence) {
		super(name,birtdate,deathdate,country);
		this.placeOfResidence = placeOfResidence;
	}
	
	
}

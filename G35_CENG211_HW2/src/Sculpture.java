
public class Sculpture extends Artworks {
	private Artist artist;
	private String material;
	private int weight;
	
	public Sculpture(String name,String style,Artist artist,String material,int weight) {
		super(name, style);
		this.artist = artist;
		this.material = material;
		this.weight = weight;
	}

}

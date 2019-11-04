
public class Painting extends Artworks{
	private Artist artist;
	private int dimension1;
	private int dimension2;
	public Painting(String name,String style,Artist artist,int dimension1,int dimension2) {
		super(name, style);
		this.artist = artist;
		this.dimension1 = dimension1;
		this.dimension2 = dimension2;
		
	}

}
